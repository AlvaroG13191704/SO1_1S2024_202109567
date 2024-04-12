package log

import (
	"context"
	"log"
	"proyecto2/consumer/models"
)

// SaveLog saves a log in the database
func (lg *LogCollection) SaveLog(value models.Log) {
	lg.Connect()

	result, err := lg.Coll.InsertOne(context.TODO(), value)
	if err != nil {
		log.Println(" Error saving log: ", err)
	}

	log.Println("Log saved: ", result)

}
