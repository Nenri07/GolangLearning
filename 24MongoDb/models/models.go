package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Youtube struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tutorial string             `json:"tutorial"`
	Watched  bool               `json:"iswatched"`
	Liked    bool               `json:"isliked"`
}
