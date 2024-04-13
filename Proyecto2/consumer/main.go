package main

import (
	"context"
	"fmt"
	"log"
	logDB "proyecto2/consumer/db/log"
	"proyecto2/consumer/models"
	"proyecto2/consumer/redis"
	"time"

	"github.com/segmentio/kafka-go"
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
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"my-cluster-kafka-bootstrap:9091", "my-cluster-kafka-bootstrap:9092", "my-cluster-kafka-bootstrap:9093"},
		Topic:     "mytopic",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// func main() {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "my-cluster-kafka-bootstrap:9092",
// 		"group.id":          "mygroupid",
// 		"auto.offset.reset": "earliest",
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to create consumer: %s\n", err)
// 		panic(err)
// 	}

// 	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	run := true

// 	for run {
// 		msg, err := c.ReadMessage(time.Second)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 		} else if !err.(kafka.Error).IsTimeout() {
// 			// The client will automatically try to recover from all errors.
// 			// Timeout is not considered an error because it is raised by
// 			// ReadMessage in absence of messages.
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 		}
// 	}

// 	c.Close()
// }

// func main() {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "my-cluster-kafka-bootstrap.so1-proyecto2.svc:9093",
// 		"group.id":          "mygroupid",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		fmt.Printf("Failed to create consumer: %s\n", err)
// 		if kafkaErr, ok := err.(kafka.Error); ok == true {
// 			fmt.Printf("Kafka error code: %s\n", kafkaErr.Code())
// 			if kafkaErr.Code() == kafka.ErrTransport {
// 				fmt.Println("Check your broker address and network settings.")
// 			}
// 		}
// 		os.Exit(1)
// 	}

// 	topic := "mytopic"
// 	err = c.SubscribeTopics([]string{topic}, nil)
// 	// Set up a channel for handling Ctrl-C, etc
// 	sigchan := make(chan os.Signal, 1)
// 	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

// 	// Process messages
// 	run := true
// 	for run {
// 		select {
// 		case sig := <-sigchan:
// 			fmt.Printf("Caught signal %v: terminating\n", sig)
// 			run = false
// 		default:
// 			ev, err := c.ReadMessage(100 * time.Millisecond)
// 			if err != nil {
// 				// Errors are informational and automatically handled by the consumer
// 				continue
// 			}
// 			fmt.Println("Message received", string(ev.Value))
// 			// fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
// 		}
// 	}

// 	c.Close()

// }

// func main() {
// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "my-cluster-kafka-brokers.so1-proyecto2.svc:9092",
// 		"group.id":          "mygroupid",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		fmt.Printf("Failed to create consumer: %s", err)
// 		os.Exit(1)
// 	}

// 	topic := "mytopic"
// 	err = c.SubscribeTopics([]string{topic}, nil)

// 	// Set up a channel for handling Ctrl-C, etc
// 	sigchan := make(chan os.Signal, 1)
// 	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

// 	// Process messages
// 	run := true
// 	for run {
// 		select {
// 		case sig := <-sigchan:
// 			fmt.Printf("Caught signal %v: terminating\n", sig)
// 			run = false
// 		default:
// 			ev, err := c.ReadMessage(100 * time.Millisecond)
// 			if err == nil {
// 				// Message successfully read
// 				var data models.Data
// 				err := json.Unmarshal(ev.Value, &data)
// 				if err != nil {
// 					fmt.Printf("Failed to unmarshal message: %s", err)
// 				} else {
// 					// Successfully unmarshalled message, now you can use the data
// 					fmt.Printf("Received data: %+v\n", data)

// 					// Process the event
// 					processEvent(data)
// 				}
// 			}
// 		}
// 	}
// }
