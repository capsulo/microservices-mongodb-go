package models

type User struct {
	ID string	`bson:"_id,omitempty"`
	Firstname string `bson:"first_name,omitempty"`
	Lastname string `bson:"last_name,omitempty"`
	Address string `bson:"address"`
}