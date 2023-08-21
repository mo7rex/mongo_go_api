package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/mo7rex/mongo_go_api/database"
	"github.com/mo7rex/mongo_go_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertOneMovie(movie models.Netflix) {
	inserted, err := database.Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the movie inserted in db with id: ", inserted.InsertedID)
}
func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) //creates a new ObjectID
	filter := bson.M{"_id": id}
	update := bson.M{"$Set": bson.M{"watched": true}}
	result, err := database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	delCount, err := database.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted count: ", delCount)
}

func deleteAllMovies() {
	delCount, err := database.Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies count deleted: ", delCount)
}

func getAllMovies() {
	cursor, err := database.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
}
