package main

import (
	"flag"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"Balancer/internal"
)

// configs
const (
	// global env TODO
	// wardenOrigin = "some ip"

	// metrics
	bindAddrHTTPMetrics = ":9001"

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
	BindAddr:     ":8088",
	ReadTimeout:  maxTimeoutWork,
	WriteTimeout: maxTimeoutWork,
	Protocol:     "http",
}

func main() {
	// App start params
	var podUUID string
	flag.StringVar(&podUUID, "pod-uuid", "1", "identification serverCfg")
	flag.Parse()

	// Metrics
	metrics := internal.NewPrometheusMetrics(serverCfg.ServiceName + "_" + podUUID)
	err := metrics.SetupMonitoring()
	if err != nil {
		logrus.Fatal(err)
	}

	// Metrics serverCfg
	go internal.CreateNewMonitoringServer(bindAddrHTTPMetrics)

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
		serverCfg.BindAddr,
		serverCfg.Protocol,
		podUUID)

	// Server
	server := internal.NewServer()
	err = server.Launch(serverCfg, routerCORS)
	if err != nil {
		logrus.Fatal(err)
	}
}
