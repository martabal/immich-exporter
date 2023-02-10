package main

import (
	"flag"
	"immich-exporter/src/immich"
	"immich-exporter/src/models"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func startup() {
	var envfile bool

	flag.BoolVar(&envfile, "e", false, "Use .env file")
	flag.Parse()
	log.Println("Loading all parameters")
	if envfile {
		useenvfile()
	} else {
		initenv()
	}

	immich.Auth()

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
