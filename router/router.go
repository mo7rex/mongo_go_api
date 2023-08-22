package router

import (
	"github.com/gorilla/mux"
	"github.com/mo7rex/mongo_go_api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetTheMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controllers.GetOne).Methods("GET")
	router.HandleFunc("/api/movies/add", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.SetWatched).Methods("PUT")
	router.HandleFunc("/api/movies/update/{id}", controllers.UpdateOne).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteAll).Methods("DELETE")

	return router
}
