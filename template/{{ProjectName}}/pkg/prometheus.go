package pkg

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Prometheus struct{}

func NewPrometheus(log *Logger, port int) *Prometheus {
	go runPrometheusHttpServer(log, port)
	return &Prometheus{}
}

func runPrometheusHttpServer(log *Logger, port int) {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Error("Failed to listen prometheus server: ", err)
	}
}
