package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	logDB "proyecto2/consumer/db/log"
	"proyecto2/consumer/models"
	"proyecto2/consumer/redis"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// TODO: What exactly does save on redis and mongo?

// Function that recieves the messages from the kafka topic and sends them to mongo and redis
func processEvent(event models.Data) {
	// Create a log object
	log := models.Log{
		Value:     event.Name,
		CreatedAt: time.Now().String(),
	}

	// save on mongo
	mongoDB := logDB.LogCollection{}
	go mongoDB.SaveLog(log)

	// save on redis
	go redis.Insert(log)
}
func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092",
		"group.id":          "mygroupid",
		"auto.offset.reset": "earliest"})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "mytopic"
	err = c.SubscribeTopics([]string{topic}, nil)

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err == nil {
				// Message successfully read
				var data models.Data
				err := json.Unmarshal(ev.Value, &data)
				if err != nil {
					fmt.Printf("Failed to unmarshal message: %s", err)
				} else {
					// Successfully unmarshalled message, now you can use the data
					fmt.Printf("Received data: %+v\n", data)

					// Process the event
					processEvent(data)
				}
			}
		}
	}
}
