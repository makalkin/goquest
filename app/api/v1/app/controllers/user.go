package controllers

import (
	"github.com/revel/revel"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	"github.com/makalkin/goquest/app/models"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)


type User struct {
	*revel.Controller
}

func (c User) Get(id string) revel.Result {
	service := services.UserService{}
	println("IDDDDDDDD", id)
	if id != "" {
		user := new(models.User)
		err := service.GetUserById(id, user)
		if (err != nil) {
			c.Response.Status = 404
			return c.RenderJson(&map[string]interface{}{"error": err.Error()})
		}

		return c.RenderJson(user)

	} else {
		var page, perPage int
		if _, ok := c.Params.Query["page"]; ok {
			page, _ = strconv.Atoi(c.Params.Query.Get("page"))
		} else {
			page = 1
		}
		if _, ok := c.Params.Query["perPage"]; ok {
			perPage, _ = strconv.Atoi(c.Params.Query.Get("perPage"))
		} else {
			perPage = 20
		}

		users, paging, err := service.GetUsers(bson.M{}, page, perPage)
		if err == nil {
			return c.RenderJson(&map[string]interface{}{"users": users, "paging": paging})
		} else {
			c.Response.Status = 400
			return c.RenderJson(&map[string]interface{}{"error": err.Error()})
		}
	}
}

func (c User) GetMe(id string) revel.Result {
	service := services.UserService{}

	user := new(models.User)
	err := service.GetUserById(id, user)
	if (err != nil) {
		c.Response.Status = 404
		return c.RenderJson(&map[string]interface{}{"error": err.Error()})
	}

	return c.RenderJson(user)
}