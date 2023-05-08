package internal

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type EchoHandler struct{}

// NewEchoHandler is constructor for EchoHandler Ping API.
func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (h *EchoHandler) Configure(r *mux.Router) {
	r.HandleFunc("/echo", h.ServeHTTP).
		Methods(http.MethodGet).
		Queries("body", "{body}")
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := bindEcho(r)
	if err != nil {
		DefaultHandlerHTTPError(r.Context(), w, err)
		return
	}

	logrus.Info(r.RemoteAddr + " send " + response.Body)

	Response(r.Context(), w, http.StatusOK, response)
}

func bindEcho(r *http.Request) (EchoResponse, error) {
	res := EchoResponse{}

	res.Body = r.FormValue("body")
	if res.Body == "" {
		return EchoResponse{}, ErrBadRequest
	}

	return res, nil
}
