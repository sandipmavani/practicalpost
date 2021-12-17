package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
// model for posts
type Posts struct {
	Id           primitive.ObjectID `json:"id" bson:"_id" validate:"nil=false"`
	Title        string             `json:"title" bson:"title"`
	Body         string             `json:"body" bson:"body"`
	ThumbnailUrl string             `json:"thumbnaiUrl" bson:"thumbnaiUrl" validate:"nil=false"`
	Author       primitive.ObjectID `json:"author" bson:"author" validate:"nil:false"` // we save auther user referance id here
	PostedOn     time.Time          `json:"postedOn" bson:"postedOn"`
}
