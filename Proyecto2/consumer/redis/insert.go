package redis

import (
	"context"
	"log"
	"proyecto2/consumer/models"
)

// Insert is a function that inserts a value
// Insert is a function that inserts a value
func Insert(value models.Log) {

	client := GetRedisInstance()

	// create or get if exists the key
	counter := int(client.Incr(context.Background(), value.Data.Name).Val())

	// insert the name as the key and the value as the counter
	err := client.HSet(context.TODO(), value.Data.Name+"_hash", "counter", counter).Err()
	if err != nil {
		log.Println("Error saving on redis: ", err)
	}

	log.Println("value saved on redis -> ", value)
}

// jsonData, err := json.Marshal(value)
// if err != nil {
// 	log.Println("Error marshalling the invitation: ", err)
// }

// err = client.Set(context.TODO(), "llave", jsonData, time.Hour*24*5).Err()
// if err != nil {
// 	log.Println("Error saving the invitation: ", err)
// }

// log.Println("value saved on redis -> ", value)
