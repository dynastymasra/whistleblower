package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/whistleblower/config"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	log := logrus.WithFields(logrus.Fields{
		"service_name": config.ServiceName,
		"version":      config.Version,
	})

	log.Infoln("Prepare start service")
}
