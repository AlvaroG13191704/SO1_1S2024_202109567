package kafka

import (
	"encoding/json"
	"fmt"
	"os"
	"proyecto2/servergRPC/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(value model.Data) interface{} {
	// create producer
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-bootstrap:9092",
		"client.id":         "goapp-producer",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	delivery_chan := make(chan kafka.Event, 10000)

	// conver to json
	jsonStr, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("Failed to serialize SMS event: %s\n", err.Error())
		return err
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(jsonStr)},
		delivery_chan,
	)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err.Error())
		os.Exit(1)
	}
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
					panic(ev.TopicPartition.Error)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
	return err
}
