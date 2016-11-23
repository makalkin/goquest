package models

import (
	"gopkg.in/mgo.v2"
)

type User struct {
	Base        `bson:",inline"`
	AccessToken string `bson:"access_token" json:"accessToken,omitempty"`
	Name        string `json:"name"`
	Fid         string `json:"fid,omitempty"`
	Circles     []struct {
		Circle     mgo.DBRef `bson:",omitempty" json:"circle"`
		Experience int       `json:"experience"`
		Currency   int       `json:"currency"`
	} `bson:",omitempty" json:"circles,omitempty"`
}
