package main

import (
	"fmt"
	"immich-exporter/src/immich"
	"immich-exporter/src/models"
	myprom "immich-exporter/src/promtheus"
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
	r := prometheus.NewRegistry()

	r.MustRegister(myprom.HttpRequestDuration)
	r.MustRegister(myprom.Version)

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	var srv *http.Server
	srv = &http.Server{Addr: ":8090", Handler: mux}
	log.Fatal(srv.ListenAndServe())
}

func metrics(w http.ResponseWriter, req *http.Request) {
	value := immich.Allrequests()
	if value == "" {
		value = immich.Allrequests()
	}

	fmt.Fprintf(w, value)
}
