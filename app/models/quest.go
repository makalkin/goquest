package models

import (
	"gopkg.in/mgo.v2"
)

type Quest struct {
	Base       `bson:",inline"`
	Circle     mgo.DBRef `bson:",omitempty" json:"circle"`
	Experience int
	Currency   int
}
