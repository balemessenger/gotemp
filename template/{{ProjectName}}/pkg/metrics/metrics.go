package metrics

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const ServiceNamespace = "{{ProjectName}}"

var (
	metricsOnce sync.Once
	metrics     *Metrics
)

type Metrics struct {
	*Measure
	MethodDurations    *prometheus.SummaryVec
	MethodCount        *prometheus.CounterVec
	MethodSuccessCount *prometheus.CounterVec
	MethodErrorCount   *prometheus.CounterVec
}

func GetMetrics() *Metrics {
	metricsOnce.Do(func() {
		metrics = NewMetrics()
		metrics.Initialize()
	})
	return metrics
}

func NewMetrics() *Metrics {
	methodLabels := []string{"service_name", "method"}
	return &Metrics{
		Measure: NewMeasure(),
		MethodDurations: prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:       "method_durations_nanoseconds",
				Help:       "Total Rpc latency.",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			}, methodLabels),
		MethodCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "method_total",
			Help: "The total number of rpc error",
		}, methodLabels),
		MethodSuccessCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "method_success_total",
			Help: "The total number of rpc error",
		}, methodLabels),
		MethodErrorCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "method_error_total",
			Help: "The total number of rpc error",
		}, methodLabels),
	}
}

func (metrics *Metrics) Initialize() {
	prometheus.MustRegister(metrics.MethodDurations)
}