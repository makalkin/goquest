package controllers

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/makalkin/goquest/app/api/v1/app/services"
	. "github.com/makalkin/goquest/app/api/v1/app/utils"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Circle struct {
	*revel.Controller
}

func (c Circle) Add() revel.Result {
	service := services.CircleService{}
	userID := bson.ObjectIdHex(c.Params.Get("userID"))
	type PostPayload struct {
		Private bool `json:"private"`
	}

	var payload PostPayload
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()})
	}

	circle := &models.Circle{
		Creator: mgo.DBRef{
			Collection: "users",
			Id:         userID,
		},
		Private: payload.Private,
	}

	err = service.Add(circle)

	if err == nil {
		return c.RenderJson(circle)
	}
	return RenderJsonError(c.Controller, 500, APIError{Msg: err.Error()})
}

func (c Circle) GetOne(id string) revel.Result {
	if !govalidator.IsMongoID(id) {
		return RenderJsonError(c.Controller, 400, APIError{Field: "id", Msg: "Not a valid mongo ID."})
	}

	messages := make(chan error, 2)
	service := services.CircleService{}
	users := services.UserService{}

	circle := new(models.Circle)
	me := new(models.User)

	go func () {
		messages <- users.GetMe(bson.M{"_id": bson.ObjectIdHex(c.Params.Get("userID"))}, me)
	}()

	go func() {
		messages <- service.Get(bson.M{"_id": bson.ObjectIdHex(id)}, circle)
	}()

	userErr, circleErr  := <- messages, <- messages
	close(messages)
	if userErr == nil && circleErr == nil {
		if circle.Private {
			for _, v := range me.Circles {
				if v.Circle.Id == circle.Id {
					return c.RenderJson(circle)
				}
			}
			return RenderJsonError(c.Controller, 400, APIError{Msg:"Circle is private."})
		}
		return c.RenderJson(circle)
	} else {
		if userErr != nil {
			return RenderJsonError(c.Controller, 400, APIError{Msg: userErr.Error()})
		}
		return RenderJsonError(c.Controller, 400, APIError{Msg: circleErr.Error()})
	}
}

func (c Circle) GetMany() revel.Result {
	return c.RenderText("")
}

func (c Circle) Update() revel.Result {
	return c.RenderText("")
}

func (c Circle) Delete() revel.Result {
	return c.RenderText("")
}

func (c Circle) Join(id string) revel.Result {
	if !govalidator.IsMongoID(id) {
		return RenderJsonError(c.Controller, 400, APIError{Field: "id", Msg: "Not a valid mongo ID."})
	}

	user := c.RenderArgs["user"].(models.User)
	userService := services.UserService{User: user}

	err := userService.AddCircle(id)
	if err != nil {
		return RenderJsonError(c.Controller, 400, APIError{Msg: err.Error()})
	}

	return c.RenderJson(true)

}

func init() {
	revel.InterceptFunc(CheckAuth, revel.BEFORE, &Circle{})
}
