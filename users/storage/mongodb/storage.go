package mongodb

import (
	"github.com/irahardianto/monorepo-microservices/users/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage struct {
	*mgo.Database
}

func (s *Storage) GetAll() []model.User {
	c := s.C("users")

	var users []model.User
	iter := c.Find(nil).Iter()
	result := model.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (s *Storage) Create(user *model.User) error {
	c := s.C("users")

	obj_id := bson.NewObjectId()
	user.Id = obj_id
	err := c.Insert(&user)
	return err
}

func (s *Storage) Delete(id string) error {
	c := s.C("users")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
