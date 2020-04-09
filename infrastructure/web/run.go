package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func Run(server *graceful.Server, router *RouterInstance) {
	logrus.WithField("address", router.port).Infoln("Start run web application")

	muxRouter := router.Router()

	server.Server = &http.Server{
		Addr: fmt.Sprintf(":%s", router.port),
		Handler: handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logrus.StandardLogger()),
		)(muxRouter),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatalln("Failed to start server")
	}
}
