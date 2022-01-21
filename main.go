package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func test() {
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "your_auth_token"

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+15558675309")
	params.SetFrom("+15017250604")
	params.SetBody("Hello from Go!")

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println("Message Sid: " + *resp.Sid)
	}
}