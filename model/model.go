package model

import (
	"context"
	"fmt"

	db "github.com/shokishimo/Distinctionality/db"
	"go.mongodb.org/mongo-driver/bson"
)

type QandA struct {
	Question string `json: question`
	Answer   string `json: answer`
}

func Get20Quizzes(version string, level string) ([]QandA, error) {
	var QandAs []QandA

	// connect to db
	client, err := db.Connect()
	if err != nil {
		return nil, err
	}
	// Disconnect from db after execution by defer
	defer db.Disconnect(client)

	// modify input string to be a correct param str to access database
	version = "Distinctionality" + version
	level = "level" + level

	// get 20 quizzes
	collection := client.Database(version).Collection(level)
	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 20}}}}}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decode one QandA data at a time
	for cursor.Next(context.TODO()) {
		var each QandA
		err := cursor.Decode(&each)
		if err != nil {
			return nil, err
		}
		QandAs = append(QandAs, each)
		fmt.Printf("%+v\n", each)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Done retireving 20 data")
	return QandAs, nil
}

func CreateQandAs(data []QandA, version string, level string) error {
	// connect to db
	client, err := db.Connect()
	if err != nil {
		return err
	}

	// begin insertOne
	collection := client.Database(version).Collection(level)

	for i := range data {
		result, err := collection.InsertOne(context.TODO(), data[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	}

	// Disconnect from db
	defer db.Disconnect(client)

	return nil
}
