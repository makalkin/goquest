package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"net/url"
	"encoding/json"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	"github.com/makalkin/goquest/app/models"
	"gopkg.in/mgo.v2/bson"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"fmt"
)


type Auth struct {
	*revel.Controller
}

func (c Auth) SignUp(token string) revel.Result {
	userData := map[string]interface{}{}
	resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/me?fields=id,name&access_token=%s", url.QueryEscape(token)))
	if err == nil {
		defer resp.Body.Close()
	}

	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		revel.ERROR.Println(err)
	} else {
		return RenderJsonError(c.Controller, 500, APIError{Msg: err.Error()})
	}

	service := services.UserService{}
	user := new(models.User)

	if err := service.GetUser(bson.M{"fid": userData["id"]}, user); err != nil {
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()}) // Look into this case
	}

	if user.IsNew() == true {
		user.Name = userData["name"].(string)
		user.Fid = userData["id"].(string)
	} else {
		return RenderJsonError(c.Controller, 400, APIError{Msg: "User is already registered."})
	}
	user.AccessToken = token

	if err := service.AddUser(user); err != nil {
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()})
	}

	c.Response.Status = 201
	return c.RenderJson("")
}