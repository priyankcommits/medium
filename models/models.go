package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Flash message Struct
	Message struct {
		Value string
	}
	// Pass keyword args to urls
	Kwargs struct {
		Key   string
		Value string
	}
	User struct {
		Nickname string
	}
	Post struct {
		ID       bson.ObjectId `bson:"_id" json:"id"`
		Nickname string        `json:"nickname"`
		Text     string        `json:"text"`
		Time     time.Time     `json:"time"`
	}
)
