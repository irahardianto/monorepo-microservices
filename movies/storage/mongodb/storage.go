package mongodb

import (
	"time"

	"github.com/irahardianto/monorepo-microservices/movies/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Storage struct {
	*mgo.Database
}

func (s *Storage) GetAll() []model.Movie {
	c := s.C("movies")

	var movies []model.Movie
	iter := c.Find(nil).Iter()
	result := model.Movie{}
	for iter.Next(&result) {
		movies = append(movies, result)
	}
	return movies
}

func (s *Storage) Create(movie *model.Movie) error {
	c := s.C("movies")

	objId := bson.NewObjectId()
	movie.Id = objId
	movie.CreatedOn = time.Now()
	err := c.Insert(&movie)
	return err
}

func (s *Storage) GetById(id string) (movie model.Movie, err error) {
	c := s.C("movies")

	err = c.FindId(bson.ObjectIdHex(id)).One(&movie)
	return
}

func (s *Storage) Delete(id string) error {
	c := s.C("movies")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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
