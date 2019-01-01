package mongodb

import (
	"github.com/irahardianto/monorepo-microservices/bookings/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage struct {
	*mgo.Database
}

func (s *Storage) GetAll() []model.Booking {
	c := s.C("bookings")

	var bookings []model.Booking
	iter := c.Find(nil).Iter()
	result := model.Booking{}
	for iter.Next(&result) {
		bookings = append(bookings, result)
	}
	return bookings
}

func (s *Storage) Create(booking *model.Booking) error {
	c := s.C("bookings")

	obj_id := bson.NewObjectId()
	booking.Id = obj_id
	err := c.Insert(&booking)
	return err
}

func (s *Storage) Delete(id string) error {
	c := s.C("bookings")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
