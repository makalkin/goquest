package models

import (
	"gopkg.in/mgo.v2"
)

type Circle struct {
	Base    `bson:",inline"`
	Creator User
	Private bool
	Quests  []mgo.DBRef
}
