package db

import (
	"assignment-2/internal/webserver/constants"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// InitializeFirestore Sets the context and client for Firestore.
func InitializeFirestore() error {
	ctx = context.Background()

	// Creates the Firebase App using the service account credentials.
	sa := option.WithCredentialsFile(constants.ServiceAccountLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return err
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	return nil

}

// CloseFirestore Closes the Firestore client.
func CloseFirestore() error {
	err := client.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetContext Returns the Firestore context.
func GetContext() context.Context {
	return ctx
}

// GetClient Returns the Firestore client.
func GetClient() *firestore.Client {
	return client
}
