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
	filter := bson.M{"username": username, "password": password}

	user := model.User{}
	err := c.Find(filter).One(&user)
	return user, err
}

func (s *Storage) StoreRefreshToken(token model.RefreshToken) error {
	c := s.C("refreshtoken")

	existingToken, err := s.GetRefreshToken(token.Token)
	if err != mgo.ErrNotFound {
		return err
	}

	if existingToken.ID != "" {
		return nil
	}

	obj_id := bson.NewObjectId()
	token.ID = obj_id
	return c.Insert(&token)
}

func (s *Storage) DeleteRefreshToken(token string) error {
	cc := s.C("authtoken")

	filter := bson.M{"token": token}
	err := cc.Remove(filter)
	return err
}

func (s *Storage) GetRefreshToken(token string) (model.RefreshToken, error) {
	c := s.C("refreshtoken")
	filter := bson.M{"token": token}

	refreshToken := model.RefreshToken{}
	err := c.Find(filter).One(&refreshToken)
	return refreshToken, err
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
