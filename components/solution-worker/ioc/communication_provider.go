package ioc

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"solution-worker/config"
	"solution-worker/internal/communication"
)

func ProvidePublisher(config config.MessageBrokerConfig) communication.Publisher {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"security.protocol": config.SecurityProtocol,
		"sasl.mechanism":    config.SaslMechanism,
		"sasl.username":     config.SaslUsername,
		"sasl.password":     config.SaslPassword,
		"ssl.ca.location":   config.SslCaLocation,
	})
	if err != nil {
		panic("Failed to create producer.")
	}
	return communication.NewKafkaPublisher(producer)
}

func ProvideSubscribe[T interface{}](config config.MessageBrokerConfig) communication.Subscriber[T] {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"group.id":          config.GroupId,
		"bootstrap.servers": config.BootstrapServers,
		"security.protocol": config.SecurityProtocol,
		"sasl.mechanism":    config.SaslMechanism,
		"sasl.username":     config.SaslUsername,
		"sasl.password":     config.SaslPassword,
		"ssl.ca.location":   config.SslCaLocation,
	})
	if err != nil {
		panic("Failed to create consumer.")
	}
	return communication.NewKafKaSubscriber[T](consumer)
}
