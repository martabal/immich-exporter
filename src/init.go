package main

import (
	"encoding/json"
	"flag"
	"immich-exporter/src/immich"
	"immich-exporter/src/models"
	"io/ioutil"
	"net/http"

	"log"
	"os"

	"github.com/joho/godotenv"
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

func startup() {
	projectinfo()
	var envfile bool

	flag.BoolVar(&envfile, "e", false, "Use .env file")
	flag.Parse()
	if envfile {
		useenvfile()
	} else {
		initenv()
	}

	immich.Auth()

}

func projectinfo() {
	fileContent, err := os.Open("./package.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var res map[string]interface{}
	json.Unmarshal([]byte(byteResult), &res)
	log.Println("Author :", res["author"])
	log.Println(res["name"], "version", res["version"])
}

func useenvfile() {
	myEnv, err := godotenv.Read()

	username := myEnv["IMMICH_USERNAME"]
	password := myEnv["IMMICH_PASSWORD"]
	immich_url := myEnv["IMMICH_BASE_URL"]
	if myEnv["IMMICH_USERNAME"] == "" {
		log.Panic("Immich username is not set.")
	}
	if myEnv["IMMICH_PASSWORD"] == "" {
		log.Panic("Immich password is not set.")
	}
	if myEnv["IMMICH_BASE_URL"] == "" {
		log.Panic("IMMICH base_url is not set.")
	}
	models.Setuser(username, password, immich_url)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Using .env file")
}

func initenv() {

	username := os.Getenv("IMMICH_USERNAME")
	password := os.Getenv("IMMICH_PASSWORD")
	immich_url := os.Getenv("IMMICH_BASE_URL")
	if os.Getenv("IMMICH_USERNAME") == "" {
		log.Panic("Immich username is not set.")

	}
	if os.Getenv("IMMICH_PASSWORD") == "" {
		log.Panic("Immich password is not set.")

	}
	if os.Getenv("IMMICH_BASE_URL") == "" {
		log.Panic("Immich base_url is not set.")

	}
	models.Setuser(username, password, immich_url)

}
