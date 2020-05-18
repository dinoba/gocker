// GOLANG factory patterns
// https://www.sohamkamani.com/blog/golang/2018-06-20-golang-factory-patterns/
package storage

import (
	"errors"
)

type Handler interface {
	Store(DockerLog) error
}

type DockerLog struct {
	Image    string `json:"image"`
	HostIP   string `json:"hostip"`
	HostName string `json:"hostname"`
	LogText  string `json:"logtext"`
	Status   string `json:"status"`
	State    string `json:"state"`
	Time     string `json:"time"`
}

//StorageNotSupported error
var StorageNotSupported = errors.New("Storage type provided is not supported...")

//GetStorageHandler storage factory function
func GetStorageHandler(storageType string, connectionString string, topic string) (Handler, error) {

	//TODO add Kafka and option to add multiple storages
	switch storageType {
	case "ELASTICSEARCH":
		return NewElasticsearchHandler(connectionString, topic)
	case "KAFKA":
		return NewKafkaHandler(connectionString, topic)
	}
	return nil, StorageNotSupported
}
