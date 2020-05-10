package main

import (
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8079", r))
}

func redisClient(s string) string {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:8081",
		DB:   0,
	})

	err := client.Set("key", s, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis value:", val)
	return val
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("make sure your api is working")
	}

	defer resp.Body.Close()
	respB, err := ioutil.ReadAll(resp.Body)

	s := redisClient(string(respB))
	fmt.Fprintf(w, "Redis Degeri: %v\n", s)
}
