package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShowTime struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Date      string             `bson:"date,omitempty"`
	CreatedAt string             `bson:"created_at,omitempty"`
	Movies    string             `bson:"movies,omitempty"`
}
