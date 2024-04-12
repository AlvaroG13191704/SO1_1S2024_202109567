package redis

import (
	"context"
	"encoding/json"
	"log"
	"proyecto2/consumer/models"
	"time"
)

// Insert is a function that inserts a value
func Insert(value models.Log) {

	client := GetRedisInstance()

	jsonData, err := json.Marshal(value)
	if err != nil {
		log.Println("Error marshalling the invitation: ", err)
	}

	err = client.Set(context.TODO(), "llave", jsonData, time.Hour*24*5).Err()
	if err != nil {
		log.Println("Error saving the invitation: ", err)
	}

	log.Println("value saved on redis -> ", value)
}
