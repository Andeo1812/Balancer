package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

type metricsI interface {
	SetupMonitoring() error
	GetRequestCounter() prometheus.Counter
	GetExecution() *prometheus.HistogramVec
}

type Middleware struct {
	cors    cors.Cors
	metrics metricsI
}

func NewMiddleware(config *CorsCfg, metrics metricsI) *Middleware {
	cors := cors.New(cors.Options{
		AllowedMethods:   config.Methods,
		AllowedOrigins:   config.Origins,
		AllowCredentials: config.Credentials,
		AllowedHeaders:   config.Headers,
		Debug:            config.Debug,
	})

	return &Middleware{
		cors:    *cors,
		metrics: metrics,
	}
}

func (m *Middleware) SetCORSMiddleware(h http.Handler) http.Handler {
	return m.cors.Handler(h)
}

func (m *Middleware) SetDefaultMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := negroni.NewResponseWriter(w)

		start := time.Now()
		h.ServeHTTP(wrappedWriter, r)

		statusCode := wrappedWriter.Status()

		m.metrics.GetExecution().
			WithLabelValues(strconv.Itoa(statusCode), r.URL.String(), r.Method).
			Observe(time.Since(start).Seconds())

		m.metrics.GetRequestCounter().Inc()
	})
}
