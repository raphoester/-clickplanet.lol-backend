package prom

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRegistry() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func HandlerForRegistry(registry *prometheus.Registry) http.Handler {
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
		Registry:          registry,
	})
}
