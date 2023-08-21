package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
