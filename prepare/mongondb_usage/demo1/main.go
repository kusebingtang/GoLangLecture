package main
import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main()  {

	var(
		client * mongo.Client
		err error
	)


	if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"));err !=nil {
		fmt.Println(err)
		return
	}

	database := client.Database("my_db")

	myCollection := database.Collection("my_collection")

	fmt.Println(myCollection.Name())

}
