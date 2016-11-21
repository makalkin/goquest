package services

import (
	. "github.com/makalkin/goquest/app/models"
	"github.com/maxwellhealth/bongo"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	user User
}

//
//func (service UserService) getUsers(query interface{}) []*User {
//	return DB.Collection("users").Find(query)
//}

func (service UserService) GetUserById(StringId string, user *User) bool {
	err := DB.Collection("users").FindById(bson.ObjectIdHex(StringId), user)
	println("GET", user, user.IsNew(), err)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		return false
	} else {
		return err == nil
	}
}

func (service UserService) GetUser(query bson.M, user *User) error {
	err := DB.Collection("users").FindOne(query, user)
	return err
}

func (service UserService) AddUser(user *User) error {
	err := DB.Collection("users").Save(user)
	return err
}

func (service UserService) UpdateUser(user *User) error {
	err := DB.Collection("users").Save(user)
	return err
}
