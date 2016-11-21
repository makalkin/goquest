package models

import "gopkg.in/mgo.v2/bson"

type Quest struct {
	Base       `bson:",inline"`
	CircleId   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Experience int
	Currency   int
}
