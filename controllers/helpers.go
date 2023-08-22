package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
func makeAsWatched(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) //creates a new ObjectID
	filter := bson.M{"_id": id}
	update := bson.M{"$Set": bson.M{"watched": true}}
	result, err := database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}
func updateMovie(movieId string, upMovie models.Netflix) {
	id, _ := primitive.ObjectIDFromHex(movieId) //creates a new ObjectID
	filter := bson.M{"_id": id}
	update := bson.M{"$Set": upMovie}
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

func deleteAllMovies() int64 {
	del, err := database.Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies count deleted: ", del.DeletedCount)
	return del.DeletedCount
}

func getAllMovies() []primitive.D {
	cursor, err := database.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.D
	for cursor.Next(context.Background()) {
		var movie bson.D
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	cursor.Close(context.Background())
	return movies
}

func getOneMovie(movieId string) primitive.D {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	res := database.Collection.FindOne(context.Background(), filter)
	var movie primitive.D
	res.Decode(&movie)
	return movie

}

//controllers

func GetTheMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	movies := getAllMovies()
	if movies == nil {
		json.NewEncoder(w).Encode("there is no movies")
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	movie := getOneMovie(params["id"])
	if movie == nil {
		json.NewEncoder(w).Encode("there is no movie with that id")
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func SetWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	makeAsWatched(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func UpdateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	var movie models.Netflix
	updateMovie(params["id"], movie)
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"] + " deleted")
}
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	dcount := deleteAllMovies()
	json.NewEncoder(w).Encode(dcount)
}
