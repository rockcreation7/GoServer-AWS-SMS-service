package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	err := SendSMS("number", "msg")
	if err != nil {
		log.Fatal(err)
	}
}

// SendSMS unit
func SendSMS(phoneNumber string, message string) error {
	AwsRegion := "us-east-1"

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(AwsRegion),
	},
	)

	svc := sns.New(sess)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(phoneNumber),
		Message:     aws.String(message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(resp) // print the response data.

	return nil
}
