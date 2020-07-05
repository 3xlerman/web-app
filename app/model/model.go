package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string             `json:"firstname" binding:"required"`
	Lastname  string             `json:"lastname" binding:"required"`
	Sex       string             `json:"sex" binding:"required"`
	Age       string             `json:"age" binding:"required"`
	Language  string             `json:"language" binding:"required"`
}
