package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type VersionMessage struct {
	Version       string
	LastCommitSha string
	Description   string
}

var VersionNumber string
var LastCommitSha string
var Message = VersionMessage{VersionNumber, LastCommitSha, "pre-interview technical test"}

func main() {
	fmt.Println("Starting http server")

	r := mux.NewRouter()
	r.HandleFunc("/version", VersionHandler)

	s := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving Version info:")
	fmt.Println(Message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Message)
}
