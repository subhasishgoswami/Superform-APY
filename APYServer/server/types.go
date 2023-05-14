package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type APYReq struct {
	UserAddress string `json:"user"`
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func loggingMiddleware(next http.Handler, logger logrus.Entry) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
