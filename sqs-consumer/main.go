package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func emailConsumer(sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		for attribute, value := range message.MessageAttributes {
			fmt.Printf("The message attributes for %s is %s", attribute, *value.StringValue)
		}
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}
	return nil
}

func main() {
	lambda.Start(emailConsumer)
}
