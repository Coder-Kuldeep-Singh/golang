package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&Person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "2", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "3", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "4", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "5", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "6", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	people = append(people, Person{ID: "7", Firstname: "Dummy", Lastname: "Singh", Address: &Address{City: "Delhi", State: "New Delhi"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
