package main

import (
	"log"
	"log-service/app"
	"net/http"
)

func main() {
	log.Println("Log service running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", app.New()))
}
