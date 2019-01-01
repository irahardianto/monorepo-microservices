package mongodb

import (
	"time"

	"github.com/irahardianto/monorepo-microservices/showtimes/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage struct {
	*mgo.Database
}

func (s *Storage) GetAll() []model.ShowTime {
	c := s.C("showtimes")

	var showtimes []model.ShowTime
	iter := c.Find(nil).Iter()
	result := model.ShowTime{}
	for iter.Next(&result) {
		showtimes = append(showtimes, result)
	}
	return showtimes
}

func (s *Storage) Create(showtime *model.ShowTime) error {
	c := s.C("showtimes")

	obj_id := bson.NewObjectId()
	showtime.Id = obj_id
	showtime.CreatedOn = time.Now()
	err := c.Insert(&showtime)
	return err
}

func (s *Storage) GetByDate(date string) (showtime model.ShowTime, err error) {
	c := s.C("showtimes")

	err = c.Find(bson.M{"date": date}).One(&showtime)
	return
}

func (s *Storage) Delete(id string) error {
	c := s.C("showtimes")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
