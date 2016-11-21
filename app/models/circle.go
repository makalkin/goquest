package models

type Circle struct {
	Base    `bson:",inline"`
	Creator User
	Private bool
}
