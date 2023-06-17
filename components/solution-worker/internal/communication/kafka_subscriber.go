package communication

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
)

type KafKaSubscriber[T interface{}] struct {
	kafkaConsumer *kafka.Consumer
}

func NewKafKaSubscriber[T interface{}](kafkaConsumer *kafka.Consumer) *KafKaSubscriber[T] {
	return &KafKaSubscriber[T]{
		kafkaConsumer: kafkaConsumer,
	}
}

func (k KafKaSubscriber[T]) Subscribe(topic string, handler func(event T) error) error {
	fmt.Printf("Subscribing on topic %s", topic)
	err := k.kafkaConsumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	run := true
	for run {
		ev := k.kafkaConsumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))
			var body T
			err := json.Unmarshal([]byte(e.Value), &body)
			if err != nil {
				return err
			}

			err = handler(body)
			if err != nil {
				fmt.Printf("%% Fail processing message  %s, because %s", string(e.Key), err.Error())
				return err
			}

			if e.Headers != nil {
				fmt.Printf("%% Headers: %v\n", e.Headers)
			}
			_, err = k.kafkaConsumer.Commit()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
					e.TopicPartition)
			}
		case kafka.Error:
			_, _ = fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	_ = k.kafkaConsumer.Close()
	return nil
}
