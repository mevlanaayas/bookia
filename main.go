package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mevlanaayas/bookia/controllers"
	u "github.com/mevlanaayas/bookia/utils"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	log.Println("App will be running on port: ", port)

	// Ping test
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		resp := u.Message(true, "success")
		u.Respond(w, resp)
	})

	router.HandleFunc("/api/me/books", controllers.GetBooksFor).Methods("GET")
	router.HandleFunc("/api/me/words", controllers.GetWordsFor).Methods("GET")
	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/words", controllers.CreateWord).Methods("POST")

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
