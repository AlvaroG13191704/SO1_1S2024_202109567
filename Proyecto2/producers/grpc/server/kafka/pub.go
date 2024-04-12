package kafka

import (
	"fmt"
	"os"
	"proyecto2/servergRPC/model"

	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(value model.Data) {
	topic := "mytopic"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Convert the Data struct to JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("Failed to marshal data: %s", err)
		os.Exit(1)
	}

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("data"),
		Value:          jsonValue,
	}, nil)

	// Wait for all messages to be delivered
	p.Flush(1 * 1000)
	p.Close()
}
