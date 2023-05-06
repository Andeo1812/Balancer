package internal

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type Metrics struct {
	executionTime *prometheus.HistogramVec
	totalHits     prometheus.Counter
}

func NewPrometheusMetrics(serviceName string) *Metrics {
	m := &Metrics{
		executionTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: serviceName + "_durations",
			Help: "DurationMinutes execution of request",
		}, []string{"status", "path", "method"}),
		totalHits: prometheus.NewCounter(prometheus.CounterOpts{
			Name: serviceName + "_total_hits",
		}),
	}

	return m
}
func (pm *Metrics) SetupMonitoring() error {
	err := prometheus.Register(pm.executionTime)
	if err != nil {
		return err
	}

	err = prometheus.Register(pm.totalHits)
	if err != nil {
		return err
	}

	return nil
}

func (pm *Metrics) GetRequestCounter() prometheus.Counter {
	return pm.totalHits
}

func (pm *Metrics) GetExecution() *prometheus.HistogramVec {
	return pm.executionTime
}

func CreateNewMonitoringServer(addr string) {
	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())

	logrus.Info("metricsI starting Server at " + addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		logrus.Fatal(err)
	}
}
