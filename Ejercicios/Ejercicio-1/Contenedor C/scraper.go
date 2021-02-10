package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "github.com/JesusAGC/dockercises/Ejercicios/Ejercicio-1/MyPackage"
	"github.com/JesusAGC/dockercises/Ejercicios/Ejercicio-1/MyPackage"
	tools "github.com/JesusAGC/dockercises/Ejercicios/Ejercicio-1/MyPackage"
)

// type Person struct {
// 	XMLName     xml.Name `xml:"person"`
// 	ID          int      `xml:"id"`
// 	FirstName   string   `xml:"first_name"`
// 	LastName    string   `xml:"last_name"`
// 	Company     string   `xml:"company"`
// 	Email       string   `xml:"email"`
// 	IPAddress   string   `xml:"ip_address"`
// 	PhoneNumber string   `xml:"phone_number"`
// }

// type ThePerson struct {
// 	ID          int
// 	FirstName   string
// 	LastName    string
// 	Company     string
// 	Email       string
// 	IPAddress   string
// 	PhoneNumber string
// }

// type Persons struct {
// 	XMLName xml.Name `xml:"people"`
// 	Persons []Person `xml:"person"`
// }

func main() {
	listP := scraper()
	// myPrinter(listP)
	databaseImport(listP)
}
func scraper() tools.Persons {

	xmlFile, err := os.Open("people.xml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Se abrio exitosamente el archivo")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var collective tools.Persons

	// fmt.Print(string(byteValue))

	xml.Unmarshal(byteValue, &collective)

	// fmt.Println(collective)

	return collective
}

func myPrinter(Pe tools.Persons) {

	for i := 0; i < len(Pe.People); i++ {
		fmt.Println(Pe.People[i].ID)
		fmt.Println("Name 1: " + Pe.People[i].FirstName)
		fmt.Println("Name 2: " + Pe.People[i].LastName)
		fmt.Println("Company: " + Pe.People[i].Company)
		fmt.Println("Email: " + Pe.People[i].Email)
		fmt.Println("IP: " + Pe.People[i].IPAddress)
		fmt.Println("Number: " + Pe.People[i].PhoneNumber)
	}

}

func databaseImport(Pe tools.Persons) {
	PersonsCollection := MyPackage.Bring_My_Collection()

	for i := 0; i < len(Pe.People); i++ {
		HumanB := tools.Person{Pe.People[i].ID, Pe.People[i].FirstName, Pe.People[i].LastName, Pe.People[i].Company, Pe.People[i].Email, Pe.People[i].IPAddress, Pe.People[i].PhoneNumber}
		insertResult, err := PersonsCollection.InsertOne(context.TODO(), HumanB)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Se inserto exitosamente el registro: ", insertResult.InsertedID)

	}

}
