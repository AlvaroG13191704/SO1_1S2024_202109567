package redis

import (
	"context"
	"log"
	"proyecto2/consumer/models"
)

// Insert is a function that inserts a value
func Insert(value models.Log) {

	client := GetRedisInstance()

	// create or get if exists the key
	counter := int(client.Incr(context.Background(), value.Data.Name).Val())

	// insert the name as the field and the value as the counter in the "counter" hash
	err := client.HSet(context.TODO(), "counter", value.Data.Name, counter).Err()
	if err != nil {
		log.Println("Error saving on redis: ", err)
	}

	log.Println("value saved on redis -> ", value)
}
