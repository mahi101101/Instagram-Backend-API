package models

import(
	"gopkg.in/mgo.v2/bson"

)

type Post struct{
	Id bson.ObjectId `json:"id" bson:"_id"`
	Caption string `json:"caption" bson:"caption"`
	ImageURL string`json:"imageurl" bson:"imageurl"`
	TimeStamp string `json:"time" bson:"time"`
	UserId string `json:"userid" bson:"userid"`
}

type Posts []Post