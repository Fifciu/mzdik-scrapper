package main

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
	"context"
	"strconv"
	"fmt"
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

func updateBusStation(busStation int, client *db.Client){
	fmt.Println("Startuje: ", busStation)
	firstWay, oppoWay := getBusStations(strconv.Itoa(busStation));

	// One way
	dataToSet := make(map[string]BusStation);
	for index, element := range firstWay {
		dataToSet[strconv.Itoa(index)] = element;
	}

	stationsRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/stations");
	ctx := context.Background();

	err := stationsRef.Set(ctx, dataToSet);
	if err != nil {
		log.Fatalln("error: ", err);
	}

	amountRef := stationsRef.Child("amount");
	err = amountRef.Set(ctx, len(dataToSet));
	if err != nil {
		log.Fatalln("error: ", err);
	}

	// Second way
	dataToSet = make(map[string]BusStation);
	for index, element := range oppoWay {
		dataToSet[strconv.Itoa(index)] = element;
	}

	stationsRef = client.NewRef("ways/backward/buses/" + strconv.Itoa(busStation) + "/stations");

	err = stationsRef.Set(ctx, dataToSet);
	if err != nil {
		log.Fatalln("error: ", err);
	}

	amountRef = stationsRef.Child("amount");
	err = amountRef.Set(ctx, len(dataToSet));
	if err != nil {
		log.Fatalln("error: ", err);
	}

	fmt.Println("Skonczylem bus: ", busStation);
}

func updateBusTimetable(busStation int, client *db.Client){
	fmt.Println("Startuje: ", busStation)

	ctx := context.Background();
	stationsRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/timetable");
	stationsBackRef := client.NewRef("ways/backward/buses/" + strconv.Itoa(busStation) + "/timetable");
	amountRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/stations/amount");
	var amount int;
	if err := amountRef.Get(ctx, &amount); err != nil {
		log.Fatalln("error: ", err);
	}

	timetable, timetableBackwards := getBusTimetable(strconv.Itoa(busStation), amount)

	err := stationsRef.Set(ctx, timetable); 
	if err != nil {
		log.Fatalln("error: ", err);
	}

	err = stationsBackRef.Set(ctx, timetableBackwards); 
	if err != nil {
		log.Fatalln("error: ", err);
	}

	/*stationsRef = client.NewRef("ways/backward/buses/" + strconv.Itoa(busStation) + "/timetable");

	err = stationsRef.Set(ctx, dataToSet);
	if err != nil {
		log.Fatalln("error: ", err);
	}

	fmt.Println("Skonczylem bus: ", busStation);*/
}