package main

import (
	"github.com/bangnh1/golang-training/07/middleware"
	"github.com/bangnh1/golang-training/07/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", routes.ReturnAllUsers).Queries("sort", "{sort}")
	myRouter.HandleFunc("/user/{id}", routes.ReturnSingleUser).Methods("GET")
	myRouter.HandleFunc("/user", routes.CreateNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", routes.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", routes.UpdateUser).Methods("PUT")
	myRouter.Use(middleware.LoggingMiddleware)

	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
