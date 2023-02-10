package main

import (
	"fmt"
	"immich-exporter/src/immich"
	"immich-exporter/src/models"

	"log"
	"net/http"
)

func main() {
	startup()
	log.Println("Immich URL :", models.GetURL())
	log.Println("username :", models.GetUsername())
	log.Println("password :", models.Getpasswordmasked())
	log.Println("Started")
	http.HandleFunc("/metrics", metrics)
	http.ListenAndServe(":8090", nil)
}

func metrics(w http.ResponseWriter, req *http.Request) {
	value := immich.Allrequests()
	if value == "" {
		value = immich.Allrequests()
	}
	fmt.Fprintf(w, value)
}
