package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	smtp "github.com/smtpcom/smtpcom-go"
	"log"
)

func main() {
	testEmail := "your_email"
	testApiKey := "your_api_key"
	testSender := "your_sender_label"

	config := smtp.NewConfiguration()
	client := smtp.NewAPIClient(config)
	ctx := context.WithValue(context.Background(), smtp.ContextAPIKey, smtp.APIKey{
		Key: testApiKey,
	})
	body := smtp.InlineObject{
		Channel: testSender, // sender label
		Recipients: smtp.V4MessagesRecipients{To: []smtp.V4MessagesRecipientsTo{
			{Name: "Test", Address: testEmail},
		}},
		Originator: smtp.V4MessagesOriginator{
			From: &smtp.V4MessagesRecipientsTo{
				Name: "Test", Address: testEmail},
		},
		CustomHeaders: nil,
		Subject:       "Test",
		Body: smtp.V4MessagesBody{Parts: []smtp.V4MessagesBodyParts{
			{Content: "Hello"},
		}},
	}
	localVarOptionals := &smtp.V4MessagesPostOpts{
		InlineObject: optional.NewInterface(body),
	}

	messageResponse, response, err := client.MessagesApi.V4MessagesPost(ctx, localVarOptionals)

	if err != nil {
		log.Println(err)
		fmt.Println(response)
	} else {
		fmt.Println(messageResponse)
		fmt.Println(response)
	}
}
