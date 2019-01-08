package mongodb

import (
	"github.com/irahardianto/monorepo-microservices/auth/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage struct {
	*mgo.Database
}

func (s *Storage) GetByUsernameAndPassword(username, password string) (model.User, error) {
	c := s.C("users")
	filter := bson.M{"username": username}

	user := model.User{}
	err := c.Find(filter).One(&user)
	return user, err
}

func (s *Storage) Ping() error {
	var pingStatus error

	err := s.Session.Ping()
	if err != nil {
		pingStatus = err
	} else {
		pingStatus = nil
	}

	return pingStatus
}
