package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting http server ")

	versionNumber := os.Getenv("VERSION")
	lastCommitSha := os.Getenv("LASTCOMMIT")
	message := VersionMessage{versionNumber, lastCommitSha, "pre-interview technical test"}

	r := mux.NewRouter()
	r.HandleFunc("/version", VersionInfo(message))

	s := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

type VersionMessage struct {
	version       string
	lastcommitsha string
	description   string
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
func VersionInfo(w http.ResponseWriter, r *http.Request, info VersionMessage) {
	json.NewEncoder(w).Encode(info)
}
