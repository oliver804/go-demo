package main

import (
	"container/list"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	MONGO_URI = "mongodb://mongo:123456@localhost:27017/test"
	DB_NAME   = "test"
)

func createCollection(collName string) {
	client := getMongoClient()
	defer closeMongoClient(client)
	client.Database(DB_NAME).CreateCollection(context.TODO(), collName)
}

func getMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("get mongo client fail, err:", err)
	}
	return client
}

func getDatas(coll string) list.List {
	client := getMongoClient()
	client.Database(DB_NAME).Collection(coll).Find(context.TODO(), mongo.NewReplaceOneModel().Filter)
}

func closeMongoClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func main() {
	var collName = "user"
	createCollection(collName)
	var users = []User{
		User{:},
	}
}
