package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event interface{}) error {
	fmt.Printf("Context: %+v\n", ctx)
	fmt.Printf("Event: %+v\n", event)
	return nil
}

func main() {
	lambda.Start(Handler)
}
