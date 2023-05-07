package internal

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Launch(config *ServerCfg, router http.Handler) error {
	ser := http.Server{
		Addr:         config.BindAddrHTTP,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}

	if config.Protocol == HTTPS {
		err := ser.ListenAndServeTLS(config.FileTLSCertificate, config.FileTLSKey)
		if err != nil {
			return fmt.Errorf("launch https: %w", err)
		}

		return nil
	}

	err := ser.ListenAndServe()
	if err != nil {
		return fmt.Errorf("launch http: %w", err)
	}

	return nil
}
