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

func dbConnect() (*db.Client, error){
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

func dbUpdateBusStation(busStation int, client *db.Client){
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

	amountRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/amount");
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

	amountRef = client.NewRef("ways/backward/buses/" + strconv.Itoa(busStation) + "/amount");
	err = amountRef.Set(ctx, len(dataToSet));
	if err != nil {
		log.Fatalln("error: ", err);
	}

	fmt.Println("Skonczylem bus: ", busStation);
}

func dbUpdateBusTimetable(busStation int, client *db.Client){
	fmt.Println("Startuje: ", busStation)

	ctx := context.Background();
	stationsRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/timetable");
	stationsBackRef := client.NewRef("ways/backward/buses/" + strconv.Itoa(busStation) + "/timetable");
	amountRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/amount");
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
}

type FewBusStations struct {
	Stations []BusStation
}

func dbGetBusStations(bus int, client *db.Client) ([]BusStation, []BusStation){
	// Remember about both side!
	ctx := context.Background();
	path := fmt.Sprintf("ways/backward/buses/%d/stations", bus);
	ref := client.NewRef(path);
	var backwardStations []BusStation;
	if err := ref.Get(ctx, &backwardStations); err != nil {
		log.Fatalln("Error reading value: ", err)
	}

	path = fmt.Sprintf("ways/forward/buses/%d/stations", bus);
	ref = client.NewRef(path);
	var forwardStations []BusStation;
	if err := ref.Get(ctx, &forwardStations); err != nil {
		log.Fatalln("Error reading value: ", err)
	}

	return backwardStations, forwardStations;
}

func dbGetBuses() []int{
	var buses []int;
	for i := 1; i <= 26; i++ {
		buses = append(buses, i)
	}
	return buses;
}

func dbGetBusCertainTimetable(bus int, backwards bool, client *db.Client, dayType int64) []BusCourse{
	// dayType 
	// 1 - CasualDay
	// 2 - Saturday
	// 3 - Saints

	var basicPath string;
	if backwards {
		basicPath = fmt.Sprintf("ways/backward/buses/%d/timetable", bus);
	} else {
		basicPath = fmt.Sprintf("ways/forward/buses/%d/timetable", bus);
	}

	criterium := "/CasualDay";
	switch(dayType){
		case 1:
			criterium = "/CasualDay"
		case 2:
			criterium = "/Saturday"
		case 3:
			criterium = "/Saints"
	}

	fullPath := fmt.Sprintf("%s%s", basicPath, criterium)
		
	ref := client.NewRef(fullPath);
	var table []BusCourse;

	ctx := context.Background();

	if err := ref.Get(ctx, &table); err != nil {
		log.Fatalln("Error reading value: ", err)
	}

	return table;
}

func dbGetBusFullTimetable(bus int, backwards bool, client *db.Client) BusTimetable{
	var path string;
	if backwards {
		path = fmt.Sprintf("ways/backward/buses/%d/timetable", bus);
	} else {
		path = fmt.Sprintf("ways/forward/buses/%d/timetable", bus);
	}

		
	ref := client.NewRef(path);
	var table BusTimetable;

	ctx := context.Background();

	if err := ref.Get(ctx, &table); err != nil {
		log.Fatalln("Error reading value: ", err)
	}

	return table;
}