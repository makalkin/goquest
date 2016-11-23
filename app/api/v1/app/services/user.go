package services

import (
	. "github.com/makalkin/goquest/app/models"
	"github.com/maxwellhealth/bongo"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	user User
}

//func (service UserService) getUsers(query bson.M)  {
//	return DB.Collection("users").Find(query)
//}

func (service UserService) GetMe(query bson.M, user *User) error {
	err := DB.Collection("users").FindOne(query, user)
	return err
}

func (service UserService) GetUserById(StringId string, user *User) error {
	err := DB.Collection("users").Collection().Find(bson.M{"_id": bson.ObjectIdHex(StringId)}).Select(bson.M{"access_token": 0, "fid": 0}).One(user)
	return err
}

func (service UserService) GetUser(query bson.M, user *User) error {
	err := DB.Collection("users").FindOne(query, user)
	return err
}

func (service UserService) GetUsers(query bson.M, page int, perPage int) ([]User, *bongo.PaginationInfo, error) {
	q := DB.Collection("users").Collection().Find(query).Select(bson.M{"access_token": 0, "fid": 0})
	rs := DB.Collection("users").Find(query)
	rs.Query = q
	paging, err := rs.Paginate(perPage, page)
	var users []User
	if err == nil {
		if rs.Error == nil {
			err := rs.Query.Iter().All(&users)
			return users, paging, err
		}
		return users, paging, rs.Error
	}
	return users, paging, err
}

func (service UserService) AddUser(user *User) error {
	err := DB.Collection("users").Save(user)
	return err
}

func (service UserService) UpdateUser(user *User) error {
	err := DB.Collection("users").Save(user)
	return err
}
