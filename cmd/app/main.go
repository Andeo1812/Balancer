package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"Balancer/internal"
)

// configs
const (
	maxTimeoutWork = 3
)

var corsCfg = &internal.CorsCfg{
	Methods:     []string{"GET"},
	Origins:     []string{"*"},
	Headers:     []string{"Content-Type", "Content-Length"},
	Credentials: false,
	Debug:       true,
}

var serverCfg = &internal.ServerCfg{
	ServiceName:  "app",
	ReadTimeout:  maxTimeoutWork,
	WriteTimeout: maxTimeoutWork,
	Protocol:     "http",
}

var metricsCfg = &internal.MetricsCfg{}

func main() {
	// Get env var
	podUUID, exists := os.LookupEnv("POD_UUID")
	if !exists {
		logrus.Fatal("Can't get POD_UUID ENV")
	}

	bindAddrHTTP, exists := os.LookupEnv("PORT_APP")
	if !exists {
		logrus.Fatal("Can't get PORT_APP ENV")
	}
	serverCfg.BindAddrHTTP = bindAddrHTTP

	bindAddrMetrics, exists := os.LookupEnv("PORT_METRICS_APP")
	if !exists {
		logrus.Fatal("Can't get PORT_METRICS_APP ENV")
	}
	metricsCfg.BindAddr = bindAddrMetrics

	metricsCfg.SourceName = serverCfg.ServiceName + "_" + podUUID + "_" + serverCfg.BindAddrHTTP

	// Metrics
	metrics := internal.NewPrometheusMetrics(metricsCfg.SourceName)
	err := metrics.SetupMonitoring()
	if err != nil {
		logrus.Fatal(err)
	}

	// Metrics serverCfg
	go internal.CreateNewMonitoringServer(metricsCfg.BindAddr)

	// Middleware
	mw := internal.NewMiddleware(corsCfg, metrics)

	// Router
	router := mux.NewRouter()

	// Api setup
	echo := internal.NewEchoHandler()
	echo.Configure(router)

	// Set middleware
	router.Use(
		mw.SetDefaultMetrics,
	)

	routerCORS := mw.SetCORSMiddleware(router)

	logrus.Infof("%s starting server at %s on protocol %s with pod uuid %s",
		serverCfg.ServiceName,
		serverCfg.BindAddrHTTP,
		serverCfg.Protocol,
		podUUID)

	// Server
	server := internal.NewServer()
	err = server.Launch(serverCfg, routerCORS)
	if err != nil {
		logrus.Fatal(err)
	}
}
