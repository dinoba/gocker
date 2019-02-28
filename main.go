package main

import (
	"github.com/dinoba/gocker/log"
)

func main() {
	//Blocking channel - script will run until any routine is alive or we will get - deadlock!
	quit := make(chan bool)

	logger := log.GetInstance()
	log.WithPrefix(logger, "Starting", "[INFO]")

	<-quit
}
