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
	HostIP   string `json:"hostip"`
	HostName string `json:"hostname"`
	Log      string `json:"log"`
	Time     string `json:"time"`
	Stream   int    `json:"stream"`
}

//StorageNotSupported error
var StorageNotSupported = errors.New("Storage type provided is not supported...")

//GetStorageHandler storage factory function
func GetStorageHandler(storageType string, connectionString string, topic string) (Handler, error) {

	//TODO add Kafka and option to add multiple storages
	switch storageType {
	case "ELASTICSEARCH":
		return NewElasticsearchHandler(connectionString, topic)
	}
	return nil, StorageNotSupported
}
