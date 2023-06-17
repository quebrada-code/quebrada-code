package communication

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaPublisher struct {
	kafkaPuducer *kafka.Producer
}

func NewKafkaPublisher(producer *kafka.Producer) *KafkaPublisher {
	return &KafkaPublisher{
		producer,
	}
}

func (k *KafkaPublisher) Send(to string, data interface{}, key string) error {
	valueBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = k.kafkaPuducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &to, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          valueBytes,
	}, nil)
	if err != nil {
		return err
	}

	// Wait for all messages to be delivered
	k.kafkaPuducer.Flush(15 * 1000)
	k.kafkaPuducer.Close()

	return nil

}
