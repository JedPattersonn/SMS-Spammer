package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {

	fmt.Println("Enter the number you would like to spam (inlcude country code): ")
	var sendNumber string
	fmt.Scanln(&sendNumber)
	fmt.Println("Enter the message in the text: ")
	var messageBody string
    	fmt.Scanln(&messageBody)
	fmt.Println("How many times would you like the message to send: ")
	var reps int
	fmt.Scanln(&reps)

    	sms(sendNumber, reps, messageBody)
}

func sms(sendNumber string, reps int, messageBody string) {
    for i := 0; i < reps; i++ {
		accountSid := "" //Add acount SID
		authToken := "" //Ad auth token
		client := twilio.NewRestClientWithParams(twilio.RestClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &openapi.CreateMessageParams{}
		params.SetTo(sendNumber) //Input number to send to with country code
		params.SetFrom("+") //Input Twilio number
		params.SetBody(messageBody) //Input message body

		resp, err := client.ApiV2010.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
			err = nil
		} else {
			fmt.Println("Send with message ID: " + *resp.Sid)
		}
    }
}
