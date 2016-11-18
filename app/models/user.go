package models

import (
	"github.com/maxwellhealth/bongo"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {

}

type User struct {
	bongo.DocumentBase `bson:",inline"`
	AccessToken string
	Name string
	Fid string
}


//
//func (service UserService) getUsers(query interface{}) []*User {
//	return DB.Collection("users").Find(query)
//}

func (service UserService) GetUserById(StringId string) (*User, bool) {
	user := &User{}
	err := DB.Collection("users").FindById(bson.ObjectIdHex(StringId), user)
	println("GET", user, user.IsNew(), err)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		return user, false
	} else {
		return user, err == nil
	}
}

func (service UserService) GetUser(query map[string]interface{}, user *User) ( error) {
	err := DB.Collection("users").FindOne(bson.M(query), user)
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
