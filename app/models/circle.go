package models

import (
	"gopkg.in/mgo.v2"
)

type Circle struct {
	Base    `bson:",inline"`
	Creator mgo.DBRef   `bson:",omitempty" json:"creator"`
	Private bool        `json:"private"`
	Quests  []mgo.DBRef `bson:",omitempty" json:"quests,omitempty"`
}
