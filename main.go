package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mevlanaayas/bookia/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	router.HandleFunc("/api/me/books", controllers.GetBooksFor).Methods("GET")
	router.HandleFunc("/api/me/words", controllers.GetWordsFor).Methods("GET")
}
