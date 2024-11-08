package config

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

	// ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	// client, _ := app.Messaging(ctx)

	return app, nil, nil
}

func SendToToken(app *firebase.App, token string, title string, body string) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
