package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Phone    int32              `json:"phone" bson:"phone"`
	UserName string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	DOB      string             `json:"dob" bson:"dob"`
}
