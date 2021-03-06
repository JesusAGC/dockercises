package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	tools "github.com/JesusAGC/dockercises/Ejercicios/Ejercicio-1/MyPackage"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

// type Person struct {
// 	ID          int
// 	FirstName   string
// 	LastName    string
// 	Company     string
// 	Email       string
// 	IPAddress   string
// 	PhoneNumber string
// }

// type Persons struct {
// 	Persons []Person
// }

//docker run -d -p 27017:27017 --name some_mongo mongo:latest
func main() {
	// BringEveryone()
	// le_search := string(BringOnePerson(5))
	// println(le_search)
	r := chi.NewRouter()

	r.Get("/people", func(w http.ResponseWriter, r *http.Request) {
		people := BringEveryone()
		json.NewEncoder(w).Encode(people)
		// json.NewEncoder(w).Encode(people)
	})

	r.Get("/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		num := chi.URLParam(r, "id")
		_id, err := strconv.ParseInt(num, 0, 64)
		if err != nil {
			print(err)
		}
		indiv := BringOnePerson(int(_id))
		json.NewEncoder(w).Encode(indiv)
		// json.NewEncoder(w).Encode(indiv)
	})

	http.ListenAndServe(":7777", r)
}

func BringOnePerson(id int) tools.Person {
	mycollection := tools.Bring_My_Collection()
	filter := bson.D{{"id", id}}
	var man tools.Person
	err := mycollection.FindOne(context.TODO(), filter).Decode(&man)

	if err != nil {
		log.Fatal(err)
	}
	// j, err := json.Marshal(man)
	// return j
	return man
}

func BringEveryone() tools.Persons {
	mycollection := tools.Bring_My_Collection()
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
	l := int(len(people))
	var collective tools.Persons
	ds := make([]tools.Person, l)
	// bsonbytes, _ := bson.Marshal(people[0])
	// bson.Unmarshal(bsonbytes, &dude)
	for i := 0; i < len(people); i++ {
		var dude tools.Person
		bsonbytes, _ := bson.Marshal(people[i])
		bson.Unmarshal(bsonbytes, &dude)
		ds[i] = dude
	}
	collective.People = ds
	// fmt.Println(collective.Persons)
	// j, err := json.Marshal(collective)
	// return j
	return collective
}

// func database() *mongo.Collection {

// 	clientOptions := options.Client().ApplyURI("mongodb://db:27017")
// 	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Conectado a MongoDB")
// 	PersonsCollection := client.Database("database_for_persons").Collection("Persons")

// 	return PersonsCollection

// }
