package main

import (
	"context"
	"follow_service/config"
	"follow_service/controllers"
	"follow_service/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	driver := config.ConnectNeo4j(ctx)
	defer driver.Close(ctx)

	followService := services.FollowService{Driver: driver}
	followController := controllers.FollowController{Service: followService}

	r := mux.NewRouter()
	r.HandleFunc("/follow", followController.FollowUser).Methods("POST")
	r.HandleFunc("/unfollow", followController.UnfollowUser).Methods("POST")
	r.HandleFunc("/followers", followController.ListFollowers).Methods("POST")
	r.HandleFunc("/following", followController.ListFollowing).Methods("POST")
	r.HandleFunc("/user", followController.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
