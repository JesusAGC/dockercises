package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	XMLName     xml.Name `xml:"person"`
	ID          int      `xml:"id"`
	FirstName   string   `xml:"first_name"`
	LastName    string   `xml:"last_name"`
	Company     string   `xml:"company"`
	Email       string   `xml:"email"`
	IPAddress   string   `xml:"ip_address"`
	PhoneNumber string   `xml:"phone_number"`
}

type ThePerson struct {
	ID          int
	FirstName   string
	LastName    string
	Company     string
	Email       string
	IPAddress   string
	PhoneNumber string
}

type Persons struct {
	XMLName xml.Name `xml:"people"`
	Persons []Person `xml:"person"`
}

func main() {
	listP := scraper()
	myPrinter(listP)
	databaseImport(listP)
}
func scraper() Persons {

	xmlFile, err := os.Open("people.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Se abrio exitosamente el archivo")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var collective Persons

	// fmt.Print(string(byteValue))

	xml.Unmarshal(byteValue, &collective)

	fmt.Println(collective)

	return collective
}

func myPrinter(Pe Persons) {

	for i := 0; i < len(Pe.Persons); i++ {
		fmt.Println(Pe.Persons[i].ID)
		fmt.Println("Name 1: " + Pe.Persons[i].FirstName)
		fmt.Println("Name 2: " + Pe.Persons[i].LastName)
		fmt.Println("Company: " + Pe.Persons[i].Company)
		fmt.Println("Email: " + Pe.Persons[i].Email)
		fmt.Println("IP: " + Pe.Persons[i].IPAddress)
		fmt.Println("Number: " + Pe.Persons[i].PhoneNumber)
	}

}

func databaseImport(Pe Persons) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a MongoDB")
	PersonsCollection := client.Database("People").Collection("Persons")

	for i := 0; i < len(Pe.Persons); i++ {
		HumanB := ThePerson{Pe.Persons[i].ID, Pe.Persons[i].FirstName, Pe.Persons[i].LastName, Pe.Persons[i].Company, Pe.Persons[i].Email, Pe.Persons[i].IPAddress, Pe.Persons[i].PhoneNumber}
		insertResult, err := PersonsCollection.InsertOne(context.TODO(), HumanB)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Se inserto exitosamente el registro: ", insertResult.InsertedID)

	}

}
