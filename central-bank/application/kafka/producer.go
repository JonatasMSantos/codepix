package kafka

import (
	"fmt"
	"os"

	kafka_ "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *kafka_.Producer {
	configMap := &kafka_.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}
	p, err := kafka_.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return p
}

func Publish(msg string, topic string, producer *kafka_.Producer, deliveryChan chan kafka_.Event) error {
	message := &kafka_.Message{
		TopicPartition: kafka_.TopicPartition{Topic: &topic, Partition: kafka_.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka_.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka_.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}
	}
}
