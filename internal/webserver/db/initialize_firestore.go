package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

func InitializeFirestore() error {
	ctx = context.Background()

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil

}

func CloseFirestore() error {
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetContext() context.Context {
	return ctx
}

func GetClient() *firestore.Client {
	return client
}
