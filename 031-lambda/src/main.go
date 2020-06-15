package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func HandleRequest(ctx context.Context, name Event) (string, error) {
	return fmt.Sprintf("Hello %s %s!", name.Name, name.Surname), nil
}
func main() {
	lambda.Start(HandleRequest)
}
