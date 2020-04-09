package middleware

import (
	"net/http"
	"time"

	"github.com/dynastymasra/whistleblower/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func Logger() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()

		next(w, r)

		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime)

		logrus.WithFields(logrus.Fields{
			"request_time":   startTime.Format(time.RFC3339),
			"delta_time":     deltaTime,
			"response_time":  responseTime.Format(time.RFC3339),
			"request_proxy":  r.RemoteAddr,
			"url":            r.URL.Path,
			"method":         r.Method,
			"request_source": r.Header.Get("X-FORWARDED-FOR"),
			"headers":        r.Header,
			config.RequestID: r.Context().Value(config.HeaderRequestID),
		}).Infoln("HTTP Request")
	}
}
