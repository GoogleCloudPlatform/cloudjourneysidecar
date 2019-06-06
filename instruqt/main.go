package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var srv = http.Server{
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
	Addr:         ":80",
	Handler:      handler(),
}

func main() {
	log.Printf("redirect server")
	srv.ListenAndServe()
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/healthz", health)
	r.HandleFunc("/", redirect)

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	re := fmt.Sprintf("%s?projectid=%s", os.Getenv("PUBLIC_URL"), os.Getenv("INSTRUCT_GCP_PROJECTID"))
	http.Redirect(w, r, re, http.StatusSeeOther)
}
