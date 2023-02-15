package main

import (
	"immich-exporter/src/immich"
	"immich-exporter/src/models"

	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func metrics(w http.ResponseWriter, r *http.Request) {
	registry := prometheus.NewRegistry()
	immich.Allrequests(registry)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)

}
