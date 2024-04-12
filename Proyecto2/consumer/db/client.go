package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

var mongoClient *mongo.Client

func Connect() *mongo.Client {

	uri := "mongodb://mongo-service:27017/" // the kubernetes service name
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func GetMongoInstance() *mongo.Client {

	if mongoClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if mongoClient == nil {
			fmt.Println("Creating single mongo instance now.")
			mongoClient = Connect()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return mongoClient
}
