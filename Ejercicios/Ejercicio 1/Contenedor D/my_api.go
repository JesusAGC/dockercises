package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Company     string
	Email       string
	IPAddress   string
	PhoneNumber string
}

type Persons struct {
	Persons []Person
}

func main() {
	BringEveryone()

}

func BringOnePerson(id int) {
	// mycollection := database()

}

func BringEveryone() {
	mycollection := database()
	ctx := context.TODO()
	cursor, err := mycollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var people []bson.M
	if err = cursor.All(ctx, &people); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(people)

	b, err := bson.MarshalExtJSON(people[0], true, true)

	fmt.Println(b)

}

func database() *mongo.Collection {

	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB")
	PersonsCollection := client.Database("database_for_persons").Collection("Persons")

	return PersonsCollection

}
