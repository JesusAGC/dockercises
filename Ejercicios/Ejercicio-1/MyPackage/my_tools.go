package MyPackage

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID          int    `xml:"id" json:"id"`
	FirstName   string `xml:"first_name" json:"first_name"`
	LastName    string `xml:"last_name" json:"last_name"`
	Company     string `xml:"company" json:"company"`
	Email       string `xml:"email" json:"email"`
	IPAddress   string `xml:"ip_address" json:"ip_address"`
	PhoneNumber string `xml:"phone_number" json:"phone_number"`
}

type Persons struct {
	People []Person `xml:"people" json:"people"`
}

func Bring_My_Collection() *mongo.Collection {

	// clientOptions := options.Client().ApplyURI("mongodb://db:27017")
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
