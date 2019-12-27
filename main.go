package main

import (
	"github.com/dinoba/gocker/log"
)

func main() {
	//Blocking channel
	quit := make(chan bool)

	logger := log.GetInstance()
	log.WithPrefix(logger, "Starting", "[INFO]")

	<-quit
}
