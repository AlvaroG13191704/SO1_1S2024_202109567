package log

import (
	"proyecto2/consumer/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogCollection struct {
	Uri    string
	Client *mongo.Client
	Coll   *mongo.Collection
}

func (sp *LogCollection) Connect() {
	sp.Client = db.GetMongoInstance()
	sp.Coll = sp.Client.Database("proyecto2").Collection("logs")
}
