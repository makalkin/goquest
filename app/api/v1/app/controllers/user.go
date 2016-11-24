package controllers

import (
	"github.com/makalkin/goquest/app/api/v1/app/services"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
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


func init() {
	revel.InterceptFunc(CheckAuth, revel.BEFORE, &User{})
}
