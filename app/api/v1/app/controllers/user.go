package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"encoding/json"
)

type User struct {
	*revel.Controller
}



func (c *User) GetOne(id string) revel.Result {
	if !govalidator.IsMongoID(id) {
		return RenderJsonError(c.Controller, 400, APIError{Field: "id", Msg: "Not a valid mongo ID."})
	}

	service := services.UserService{}
	user := new(models.User)
	err := service.GetUserById(id, user)
	if err != nil {
		return RenderJsonError(c.Controller, 404, APIError{Msg: err.Error()})
	}
	return c.RenderJson(user)
}

func (c *User) GetMany() revel.Result {
	service := services.UserService{}
	revel.INFO.Println("WTF", c.Params.Query.Get("userId"))
	page, perPage := GetPaging(c.Controller)

	users, paging, err := service.GetUsers(nil, page, perPage)
	if err == nil {
		return c.RenderJson(&map[string]interface{}{"users": users, "paging": paging})
	} else {
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()})
	}
}

func (c User) GetMe() revel.Result {
	service := services.UserService{}
	user := new(models.User)
	err := service.GetMe(bson.M{"_id": bson.ObjectIdHex(c.Params.Query.Get("userId"))}, user)
	if err != nil {
		return RenderJsonError(c.Controller, 404, APIError{Msg: err.Error()})
	}

	return c.RenderJson(user)
}

func (c User) Add(token string) revel.Result {
	userData := map[string]interface{}{}
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name&access_token=" +
		url.QueryEscape(token))
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
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()})	// Look into this case
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

	return c.RenderJson("")
}


func init() {
	revel.InterceptFunc(CheckAuth, revel.BEFORE, &User{})
}
