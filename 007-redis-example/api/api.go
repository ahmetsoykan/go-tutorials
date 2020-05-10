package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID      string
	Name    string
	Surname string
}

type Box struct {
	Items []User
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp1 := User{
		ID:      "101",
		Name:    "Ahmet",
		Surname: "Soykan",
	}
	resp2 := User{
		ID:      "102",
		Name:    "Selcuk",
		Surname: "Soykan",
	}

	box := Box{}
	box.Items = append(box.Items, resp1)
	box.Items = append(box.Items, resp2)
	resp, _ := json.Marshal(box.Items)
	fmt.Fprint(w, string(resp))
}
