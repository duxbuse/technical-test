package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type VersionMessage struct {
	version       string
	lastcommitsha string
	description   string
}

var versionNumber = os.Getenv("VERSION")
var lastCommitSha = os.Getenv("LASTCOMMIT")
var message = VersionMessage{versionNumber, lastCommitSha, "pre-interview technical test"}

func main() {
	fmt.Println("starting http server ")

	r := mux.NewRouter()
	r.HandleFunc("/version", VersionHandler)

	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

/*
--- EXAMPLE OUPUT ---
"myapplication": [
  {
"version": "1.0",
"lastcommitsha": "abc57858585",
"description" : "pre-interview technical test"
} ]
*/
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
