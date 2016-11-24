package controllers

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	*revel.Controller
}

type APIError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func (c *User) GetOne(id string) revel.Result {
	if !govalidator.IsMongoID(id) {
		return RenderJsonError(c.Controller, 400, APIError{Field: "id", Msg: "Not a valid mongo ID."})
	}

	service := services.UserService{}
	user := new(models.User)
	err := service.GetUserById(id, user)
	if err != nil {
		return RenderJsonError(c.Controller, 404, err)
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

func init() {
	revel.InterceptFunc(CheckAuth, revel.BEFORE, &User{})
}
