package utils

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"strings"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"fmt"
	"net/url"
	"encoding/json"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	"errors"
)


func CheckAuth(c *revel.Controller) revel.Result {
	// Check for Auth token. Must be a valid facebook token from client.
	var token string
	tokenHeader := c.Request.Header.Get("Authorization")
	tokenParts := strings.Split(tokenHeader, "Bearer ")
	if len(tokenParts) == 2 {
		token = string(tokenParts[1])
	} else {
		return RenderJsonError(c, 401, errors.New("Invalid auth header."))
	}

	if token == "" {
		return RenderJsonError(c, 401, errors.New("Authentication token is missing."))
	}
	// Try to retrieve userID from redis.
	userId, err := redis.String(models.Redis.Do("GET", token))

	if err == nil && userId != "" {
		// User exists so we let request through with retrieved user ID.
		c.Params.Query.Set("userId", userId)
	} else {
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

		if err == nil {
			defer resp.Body.Close()
		} else {
			RenderJsonError(c, 500, errors.New("Failed to fetch user from facebook."))
		}

		if err := json.NewDecoder(resp.Body).Decode(&tokenData); err != nil {
			RenderJsonError(c, 500, errors.New("Failed to decode facebook response."))
		}

		service := services.UserService{}
		me := &models.User{}
		tokenData = tokenData["data"].(map[string]interface{})
		if isValid, ok := tokenData["is_valid"].(bool); ok && isValid {
			// Now that we have fid we can get user from our DB.
			err = service.GetMe(bson.M{"fid": tokenData["user_id"]}, me)
			if err == nil {
				if me.IsNew() == true {
					return RenderJsonError(
						c,
						401,
						errors.New("This user is not registered please authenticate first."),
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
				c.Params.Query.Set("userId", me.Id.Hex())
			} else {
				return RenderJsonError(c, 500, err)
			}
		} else {
			return RenderJsonError(c, 401, errors.New("Invalid token."))
		}
	}
	// Everything is ok. We've got the user. Move on.
	return nil
}

