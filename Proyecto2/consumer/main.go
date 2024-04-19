package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	logDB "proyecto2/consumer/db/log"
	"proyecto2/consumer/models"
	"proyecto2/consumer/redis"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

// TODO: What exactly does save on redis and mongo?

// Function that recieves the messages from the kafka topic and sends them to mongo and redis
func processEvent(event []byte) {

	// unmarshal the data
	var data models.Data
	err := json.Unmarshal(event, &data)
	if err != nil {
		fmt.Printf("Failed to unmarshal message: %s", err)
		return
	}

	// Create a log object
	log := models.Log{
		Data:      data,
		CreatedAt: time.Now().String(),
	}

	// save on mongo
	mongoDB := logDB.LogCollection{}
	go mongoDB.SaveLog(log)

	// save on redis
	go redis.Insert(log)
}

func main() {
	// to consume messages
	topic := "mytopic"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"my-cluster-kafka-bootstrap:9092"},
		Topic:       topic,
		Partition:   0,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
		GroupID:     uuid.New().String(), // Generate a unique GroupID
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("failed to read message:", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		// Process the event
		processEvent(m.Value)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
