package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("D:\\go\\path\\src\\git.marykay.com.cn\\study\\push\\cmd\\serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "cKRGCbntBak:APA91bGjYLdwx4eEoJtRU8s-goXD5E0vQ2ATto76kqcopRlOHaMmR6a1xOAZzwhHOhG4xR2zfbDmdays3OIeQPw9O923D07ve4L9nKefP8XHdodH5P2nx7n5U4YCkK2ph9szjp6imllb"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}
