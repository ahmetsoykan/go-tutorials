package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, urlData Event) (string, error) {

	//fmt.Print(name.Text)
	data := []byte(urlData.Text)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://993655.execute-api.eu-west-1.amazonaws.com/", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
func main() {
	lambda.Start(HandleRequest)
}
