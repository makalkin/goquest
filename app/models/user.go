package models

import (
	"gopkg.in/mgo.v2"
)

type UserCircle struct {
	Circle     mgo.DBRef `bson:",omitempty" json:"circle"`
	Experience int       `json:"experience"`
	Currency   int       `json:"currency"`
}

type User struct {
	Base        `bson:",inline"`
	AccessToken string       `bson:"access_token" json:"accessToken,omitempty"`
	Name        string       `json:"name"`
	Fid         string       `json:"fid,omitempty"`
	Circles     []UserCircle `bson:",omitempty" json:"circles,omitempty"`
}
