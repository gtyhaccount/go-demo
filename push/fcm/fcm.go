package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("D:\\go\\path\\src\\git.marykay.com.cn\\study\\push\\fcm\\elearning-tw.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "cwD2CULe3ZM:APA91bHcIlxbF6Qq0WxinrvJJ52_4gg7PGq3hFLTie2mxZlX6xevbD-Ymwx0YqtbQND6IdhmwzmXsCMISfh0KtW7ZrdWS56PHBEf7gbTobgyayEXNnibI1Y6iGoa1faQflncX66MG2Wa"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	//// Send a message in the dry run mode.
	//response, err := client.SendDryRun(ctx, message)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//// Response is a message ID string.
	//fmt.Println("Dry run successful:", response)

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
