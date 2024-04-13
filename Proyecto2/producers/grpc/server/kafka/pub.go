package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"proyecto2/servergRPC/model"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(value model.Data) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       &tls.Config{}, // Enable TLS
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"my-cluster-kafka-bootstrap:9091", "my-cluster-kafka-bootstrap:9092", "my-cluster-kafka-bootstrap:9093"},
		Topic:   "mytopic",
		Dialer:  dialer, // Use the custom dialer
	})

	// Convert the data struct into a byte slice
	valueBytes, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("Failed to marshal value: %v", err)
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: valueBytes,
	})

	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
	}
}

// func Produce(value model.Data) {
// 	topic := "mytopic"
// 	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092"})
// 	if err != nil {
// 		fmt.Printf("Failed to create producer: %s", err)
// 		os.Exit(1)
// 	}

// 	// Convert the Data struct to JSON
// 	jsonValue, err := json.Marshal(value)
// 	if err != nil {
// 		fmt.Printf("Failed to marshal data: %s", err)
// 		os.Exit(1)
// 	}

// 	p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 		Key:            []byte("data"),
// 		Value:          jsonValue,
// 	}, nil)

// 	// Wait for all messages to be delivered
// 	p.Flush(1 * 1000)
// 	p.Close()
// }
