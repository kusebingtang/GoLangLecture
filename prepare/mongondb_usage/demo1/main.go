package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {

	var (
		client *mongo.Client
		err    error
	)

	if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(time.Second * 5)); err != nil {
		fmt.Println(err)
		return
	}

	database := client.Database("my_db")

	myCollection := database.Collection("my_collection")

	fmt.Println(myCollection.Name())

}
