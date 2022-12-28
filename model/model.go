package model

import (
	"context"
	"fmt"

	db "github.com/shokishimo/Distinctionality/db"
)

type QandA struct {
	Question string `json: question`
	Answer   string `json: answer`
}

func getAllQandA() []QandA {
	var QandAs []QandA

	return QandAs
}

func CreateQandAs(data []QandA) error {
	// connect to db
	client, err := db.Connect()
	if err != nil {
		return err
	}

	// begin insertOne
	collection := client.Database("DistinctionalityCluster").Collection("QandA")

	for i := range data {
		result, err := collection.InsertOne(context.TODO(), data[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	}

	// Disconnect from db
	db.Disconnect(client)

	return nil
}
