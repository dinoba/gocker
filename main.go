package main

import (
	"github.com/dinoba/gocker/collector"
	config "github.com/dinoba/gocker/customconfig"
	"github.com/dinoba/gocker/log"
	"github.com/dinoba/gocker/storage"
)

func main() {
	//Blocking channel
	quit := make(chan bool)

	logger := log.GetInstance()
	log.WithPrefix(logger, "Starting", "[INFO]")
	conf, err := config.GetInstance()
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	//read config file
	elasticsearchhost := conf.GetString("storage.elasticsearchhost")
	//kafkahost := conf.GetString("storage.kafkahost")
	topic := conf.GetString("storage.topic") //kafka topic =  Elasticsearch index
	dockersToSkip := conf.GetString("skip.dockers_to_skip")

	st, err := storage.GetStorageHandler("ELASTICSEARCH", elasticsearchhost, topic)
	//st, err := storage.GetStorageHandler("KAFKA", kafkahost, topic)

	col := collector.NewCollector(st, dockersToSkip)
	err = col.CollectStats()
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}
	<-quit
}
