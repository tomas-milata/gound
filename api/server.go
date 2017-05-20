package api

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/command", CommandHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
