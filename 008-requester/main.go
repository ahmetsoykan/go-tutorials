package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"firstname": "Nic", "lastname": "Rabby"}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	// istek atma metodu 1
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, _ = client.Do(request)
	// -- end --
	//istek atma metodu 2
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	// -- end --
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
