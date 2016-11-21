package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Base        `bson:",inline"`
	AccessToken string
	Name        string
	Fid         string
	Circles     []struct {
		CircleId   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Experience int
		Currency   int
	} `bson:",omitempty"`
}
