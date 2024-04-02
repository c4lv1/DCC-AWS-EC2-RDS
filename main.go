package main

import (
	"fmt"
	"net/http"
)

type user struct {
	name string
	age  int
}

var users = []user{
	{name: "Alice", age: 25},
	{name: "Ali", age: 25},
	{name: "Alic", age: 29},
	{name: "Alil", age: 27},
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
}
func Users(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(users); i++ {
		fmt.Fprintf(w, "name: %v \n age: %v\n\n", users[i].name, users[i].age)
	}
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users", Users)

	http.ListenAndServe(":8080", nil)
}
