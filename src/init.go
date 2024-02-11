package main

import (
	"flag"
	"fmt"
	"log/slog"

	"immich-exp/immich"
	"immich-exp/logger"
	"immich-exp/models"

	"net/http"
	"strconv"
	"strings"

	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const DEFAULTPORT = 8090

var (
	Version     = "dev"
	Author      = "martabal"
	ProjectName = "immich-exporter"
)

func main() {
	loadenv()

	fmt.Printf("%s (version %s)\n", ProjectName, Version)
	fmt.Println("Author:", Author)
	fmt.Println("Using log level: ", models.GetLogLevel())
	logger.Log.Info("Immich URL: " + models.Getbaseurl())
	logger.Log.Info("Started")
	http.HandleFunc("/metrics", metrics)
	addr := ":" + strconv.Itoa(models.GetPort())
	logger.Log.Info("Listening on port " + strconv.Itoa(models.GetPort()))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err.Error())
	}
}

func metrics(w http.ResponseWriter, r *http.Request) {
	logger.Log.Debug("New request")
	registry := prometheus.NewRegistry()
	immich.Allrequests(registry)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func loadenv() {
	var envfile bool
	flag.BoolVar(&envfile, "e", false, "Use .env file")
	flag.Parse()
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) && !envfile {
		err := godotenv.Load(".env")
		if err != nil {
			errormessage := "Error loading .env file:" + err.Error()
			panic(errormessage)
		}

	}

	immichapikey := getEnv("IMMICH_API_KEY", "", false, "Immich API Key is not set", true)
	immichURL := getEnv("IMMICH_BASE_URL", "http://localhost:8080", true, "Qbittorrent base_url is not set. Using default base_url", false)
	exporterPort := getEnv("EXPORTER_PORT", strconv.Itoa(DEFAULTPORT), false, "", false)

	num, err := strconv.Atoi(exporterPort)

	if err != nil {
		panic("EXPORTER_PORT must be an integer")
	}
	if num < 0 || num > 65353 {
		panic("EXPORTER_PORT must be > 0 and < 65353")
	}
	loglevel := setLogLevel(getEnv("LOG_LEVEL", "INFO", false, "", false))

	models.SetApp(num, false, loglevel)
	models.Setuser(immichURL, immichapikey)
}
func setLogLevel(logLevel string) string {
	upperLogLevel := strings.ToUpper(logLevel)
	level, found := logger.LogLevels[upperLogLevel]
	if !found {
		upperLogLevel = "INFO"
		level = logger.LogLevels[upperLogLevel]
	}

	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.Level(level),
		},
	}
	handler := logger.NewPrettyHandler(os.Stdout, opts)
	logger.Log = slog.New(handler)
	return upperLogLevel
}

func getEnv(key string, fallback string, printLog bool, logPrinted string, needed bool) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		if needed {
			errormessage := "Please set a value for " + key
			panic(errormessage)
		}
		if printLog {
			logger.Log.Warn(logPrinted)
		}
		return fallback
	}
	return value
}
