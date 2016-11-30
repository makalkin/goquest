package services

import (
	. "github.com/makalkin/goquest/app/models"
	"gopkg.in/mgo.v2/bson"
)

type CircleService struct {
}

func (s *CircleService) Add(circle *Circle) error {
	err := DB.Collection("circles").Save(circle)
	return err
}

func (s *CircleService) Get(query bson.M, circle *Circle) error {
	err := DB.Collection("circles").FindOne(query, circle)
	return err
}