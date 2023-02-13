package models

import "github.com/prometheus/client_golang/prometheus"

var (
	appVersion string
	version    = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": appVersion,
		},
	})

	pubhttpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "method"})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of all HTTP requests",
	}, []string{"code", "handler", "method"})
)
