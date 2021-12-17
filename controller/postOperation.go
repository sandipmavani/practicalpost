package controller

import (
	"log"
	"practicalpost/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// here we perform create opertion of post
func CreatePost(post *models.Posts) (string, error) {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	post.Id = primitive.NewObjectID()

	result, err := client.Database("post_db").Collection("posts").InsertOne(ctx, post)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't create the post")
	}
	uid := result.InsertedID.(primitive.ObjectID).Hex()
	return uid, err
}

// here we perform edit opertion of post
func EditPost(postId string, post *models.Posts) (int64, error) {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	id, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.M{"_id": id}
	update := bson.D{{"$set", bson.D{
		{"title", post.Title},
		{"body", post.Body},
		{"thumbnaiUrl", post.ThumbnailUrl},
	}}}
	// we only update title,body, thumbanil field
	result, err := client.Database("post_db").Collection("posts").UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't update the post")
	}
	updateCount := result.ModifiedCount
	return updateCount, err
}

//get all post
func GetAllPost() ([]models.Posts, error) {
	var postList []models.Posts
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	findOptions := options.Find()
	// Sort by `postedOn` field descending
	findOptions.SetSort(bson.D{{"postedOn", -1}})

	cursor, err := client.Database("post_db").Collection("posts").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("Coun't get the Post")
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &postList); err != nil {

	}
	return postList, err
}

// here we perform delete opertion of post
func DeletePost(postId string) (int64, error) {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	id, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.M{"_id": id}

	// we only update title,body, thumbanil field
	result, err := client.Database("post_db").Collection("posts").DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		log.Printf("Couldn't update the post")
	}
	deleteCount := result.DeletedCount
	return deleteCount, err
}
