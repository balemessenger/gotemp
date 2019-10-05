package metrics

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "{{ProjectName}}"

var (
	metricsOnce sync.Once
	metrics     *Metrics
)

type Metrics struct {
	*Measure
	TotalPushDurations prometheus.Summary
	ErrorCount         prometheus.Counter
}

func GetMetrics() *Metrics {
	metricsOnce.Do(func() {
		metrics = NewMetrics()
		metrics.Initialize()
	})
	return metrics
}

func NewMetrics() *Metrics {
	return &Metrics{
		Measure: NewMeasure(),
		TotalPushDurations: prometheus.NewSummary(
			prometheus.SummaryOpts{
				Namespace:  namespace,
				Name:       "total_durations_nanoseconds",
				Subsystem:  "push",
				Help:       "Total Push latency.",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
		),
		ErrorCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "error_ops",
			Help:      "The total number of error",
		}),
	}
}

func (metrics *Metrics) Initialize() {
	fmt.Println("Registering Metrics...")
	prometheus.MustRegister(metrics.TotalPushDurations)
}
