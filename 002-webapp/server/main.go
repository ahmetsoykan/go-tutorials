package main

import (
	"log"
	"net/http"

	"github.com/ahmetsoykan/go-tutorials/002-webapp/gotr/client/server/api/user"
)

func main() {
	// register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
