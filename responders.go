package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/appengine/log"
)

func sendJSON(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json;  charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if content == "null" || content == "[]" {
		w.WriteHeader(404)
		content = "{ \"error\" : \"Not Found\" }"
	}

	fmt.Fprint(w, content)
}

func sendMessage(w http.ResponseWriter, msg string) {
	content := "{ \"msg\" : \"" + msg + "\" }"
	sendJSON(w, content)
}

func handleError(c context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json;  charset=UTF-8")
	w.WriteHeader(500)
	content := "{ \"error\" : \"" + err.Error() + "\" }"
	sendJSON(w, content)

	log.Errorf(c, err.Error())
}
