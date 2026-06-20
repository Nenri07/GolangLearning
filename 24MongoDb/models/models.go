package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Youtube struct {
	Id       *bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tutorial string         `json:"tutorial"`
	Watched  bool           `json:"iswatched"`
	Liked    bool           `json:"isliked"`
}
