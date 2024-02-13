package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Title  string             `bson:"title,omitempty" json:"title,omitempty" validate:"required"`
	Author string             `bson:"author,omitempty" json:"author,omitempty" validate:"required"`
	Genre  string             `bson:"genre,omitempty" json:"genre,omitempty" validate:"required"`
}
