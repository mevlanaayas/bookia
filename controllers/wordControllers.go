package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mevlanaayas/bookia/models"
	u "github.com/mevlanaayas/bookia/utils"
	"net/http"
	"strconv"
)

var CreateWord = func(w http.ResponseWriter, r *http.Request) {
	word := &models.Word{}

	err := json.NewDecoder(r.Body).Decode(word)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := word.Create()
	u.Respond(w, resp)
}

var GetWordsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	strBookId := params["bookId"]
	bookId, _ := strconv.Atoi(strBookId)

	// select by if condition
	data := models.GetWordsByBook(username, bookId)
	data := models.GetWords(username)

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
