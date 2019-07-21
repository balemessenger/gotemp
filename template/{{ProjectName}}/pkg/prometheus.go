package pkg

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	promOnce sync.Once
	prom     *Prometheus
)

type Prometheus struct {
}

func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

func GetPrometheus() *Prometheus {
	promOnce.Do(func() {
		prom = NewPrometheus()
	})
	return prom
}

func (p Prometheus) Initialize(port int) {
	go p.runPrometheusHttpServer(port)
}

func (Prometheus) runPrometheusHttpServer(port int) {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		GetLog().Error("Failed to listen prometheus server: ", err)
	}
}
