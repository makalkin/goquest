package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"strings"
)

func CheckAuth(c *revel.Controller) revel.Result {
	// Check for Auth token. Must be a valid facebook token from client.
	tokenHeader := c.Request.Header.Get("Authorization")
	tokenParts := strings.Split(tokenHeader, "Bearer ")

	var token string

	if len(tokenParts) != 2 {
		return RenderJsonError(c, 401, errors.New("Invalid auth header."))
	}

	token = string(tokenParts[1])

	if token == "" {
		return RenderJsonError(c, 401, errors.New("Authentication token is missing."))
	}

	// Try to retrieve userID from redis.
	userID, err := redis.String(models.Redis.Do("GET", token))

	if err == nil && userID != "" {
		// User exists so we let request through with retrieved user ID.
		c.Params.Set("userID", userID)

		return nil
	}

	// Token wasn't verified so we try to debug it.
	tokenData := map[string]interface{}{}
	resp, err := http.Get(
		fmt.Sprintf("https://graph.facebook.com/debug_token?input_token=%s&access_token=%s",
			url.QueryEscape(token),
			url.QueryEscape(
				fmt.Sprintf("%s|%s", models.FACEBOOK.ClientID, models.FACEBOOK.ClientSecret),
			),
		),
	)

	if err != nil {
		RenderJsonError(c, 500, errors.New("Failed to fetch user from facebook."))
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&tokenData); err != nil {
		RenderJsonError(c, 500, errors.New("Failed to decode facebook response."))
	}

	service := services.UserService{}
	me := &models.User{}
	tokenData = tokenData["data"].(map[string]interface{})

	if isValid, ok := tokenData["is_valid"].(bool); !ok || !isValid {
		return RenderJsonError(c, 401, errors.New("Invalid token."))
	}

	// Now that we have fid we can get user from our DB.
	err = service.GetMe(bson.M{"fid": tokenData["user_id"]}, me)

	if err != nil {
		return RenderJsonError(c, 500, err)
	}

	if me.IsNew() == true {
		return RenderJsonError(
			c,
			401,
			errors.New("This user is not registered please sign up first."),
		)
	}

	// Update token if it was refreshed.
	if me.AccessToken != token {
		// Remove old token.
		_, err = models.Redis.Do("DEL", me.AccessToken)
		me.AccessToken = token
		err = service.UpdateUser(me)

		if err != nil {
			return RenderJsonError(c, 500, err)
		}
	}

	_, err = models.Redis.Do("SET", token, me.Id.Hex())
	c.Params.Set("userID", me.Id.Hex())

	// Everything is ok. We've got the user. Move on.
	return nil
}
