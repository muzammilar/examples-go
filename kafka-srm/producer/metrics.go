// The module producer contains the code for a sample data producer

package main

import (
	"net/http"
	"time"

	"github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rcrowley/go-metrics"
)

// metrics registry
var MetricsRegistry metrics.Registry

func startMetricsCollector() {
	// initialize a new metrics registry
	MetricsRegistry = metrics.NewRegistry() // alternatively we can use metrics.DefaultRegistry
	// setup the prometheus client for go-metrics
	prometheusClient := prometheusmetrics.NewPrometheusProvider(
		MetricsRegistry, "kafkasrm", "producer", prometheus.DefaultRegisterer, 1*time.Second)
	// start metrics handler
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()
	// start periodic updater
	go prometheusClient.UpdatePrometheusMetrics()

}
