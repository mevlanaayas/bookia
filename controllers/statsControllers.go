package controllers

import (
	"github.com/mevlanaayas/bookia/models"
	u "github.com/mevlanaayas/bookia/utils"
	"net/http"
)

var GetStats = func(w http.ResponseWriter, r *http.Request) {
	stats := &u.Stats{}
	stats.WordCount = models.WordCount()
	stats.BookCount = models.BookCount()
	resp := u.Message(true, "success")
	resp["data"] = stats
	u.Respond(w, resp)
}
