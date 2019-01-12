package main

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
	"context"
)

func connect() (*db.Client, error){
	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://mzdik-scrapper.firebaseio.com",
	  }
	opt := option.WithCredentialsFile("../mzdik-scrapper.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return client, err;
}