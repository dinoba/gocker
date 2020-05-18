package storage

import (
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
)

//SendMessageToKafka desc
func SendMessageToKafka(brokerList []string, topic string, message string) {

}

//KafkaHandler ..
type KafkaHandler struct {
	producer sarama.AsyncProducer
	topic    string
}

//NewKafkaHandler creates kafka handler
func NewKafkaHandler(kafkaServer string, topic string) (*KafkaHandler, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionNone

	pr, err := sarama.NewAsyncProducer([]string{kafkaServer}, config)

	return &KafkaHandler{
			topic:    topic,
			producer: pr,
		},
		err
}

//StoreLog send log to kafka
func (handler *KafkaHandler) Store(log DockerLog) (err error) {
	logBytes, err := json.Marshal(log)

	handler.producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     handler.topic,
		Timestamp: time.Now(),
		Value:     sarama.ByteEncoder(logBytes),
	}

	return
}
