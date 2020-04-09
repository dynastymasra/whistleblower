package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dynastymasra/whistleblower/config"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
}
