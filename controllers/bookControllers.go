package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mevlanaayas/bookia/models"
	u "github.com/mevlanaayas/bookia/utils"
	"net/http"
)

var CreateBook = func(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := book.Create()
	u.Respond(w, resp)
}

var GetBooksFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	username := params["username"]
	data := models.GetBooks(username)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
