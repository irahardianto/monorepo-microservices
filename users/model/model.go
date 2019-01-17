package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string        `json:"firstname"`
		LastName  string        `json:"lastname"`
		Username  string        `json:"username"`
		Password  string        `json:"password"`
	}
)
