package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"strings"
	"github.com/garyburd/redigo/redis"
)

type User struct {
	*revel.Controller
}

func (c *User) GetOne(id string) revel.Result {
	service := services.UserService{}
	user := new(models.User)
	err := service.GetUserById(id, user)
	if err != nil {
		return RenderJsonError(c.Controller, 404, err)
	}
	return c.RenderJson(user)
}

func (c *User) GetMany(id string) revel.Result {
	service := services.UserService{}
	revel.INFO.Println("WTF", c.Params.Query.Get("userId"))
	page, perPage := GetPaging(c.Controller)


	users, paging, err := service.GetUsers(nil, page, perPage)
	if err == nil {
		return c.RenderJson(&map[string]interface{}{"users": users, "paging": paging})
	} else {
		return RenderJsonError(c.Controller, 400, err)
	}
}

func (c User) GetMe() revel.Result {
	service := services.UserService{}
	user := new(models.User)
	err := service.GetMe(bson.M{"_id": bson.ObjectIdHex(c.Params.Query.Get("userId"))}, user)
	if err != nil {
		return RenderJsonError(c.Controller, 404, err)
	}

	return c.RenderJson(user)
}

func checkAuth(c *revel.Controller) revel.Result {
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
	// Try to retrieve userID from redis

	userId, err := redis.String(models.Redis.Do("GET", token))

	revel.INFO.Println("TOKEN DATA",userId, err)

	if err == nil && userId != "" {
		// User exists so we let request through with retrieved user ID.
		c.Params.Query.Set("userId", userId)
	} else {
		// Token wasn't verified so we try to debug it.
		tokenData := map[string]interface{}{}
		revel.INFO.Println("WTF 1")
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
		revel.INFO.Println("WTF 1", tokenData["data"])
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

func init() {
	revel.InterceptFunc(checkAuth, revel.BEFORE, &User{})
}
