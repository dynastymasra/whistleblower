package main

import "github.com/dynastymasra/whistleblower/config"

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {

}
