package server

import (
	"github.com/mdalbrid/utils/logger"
	"net/http"
)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		logger.Error("Ping handler - response write error ", err)
	}
}
