package main

import (
	"context"
	"log"
	"firebase.google.com/go/db"
	"strconv"
	"fmt"
)

func updateBusStation(mainCtx context.Context, busStation int, client *db.Client){
	fmt.Println("Startuje: ", busStation)
	firstWay, oppoWay := getBusStations(strconv.Itoa(busStation));

	// One way
	dataToSet := make(map[string]BusStation);
	for index, element := range firstWay {
		dataToSet[strconv.Itoa(index)] = element;
	}

	stationsRef := client.NewRef("ways/forward/buses/" + strconv.Itoa(busStation) + "/stations");
	ctx, cancelFunc := context.WithCancel(mainCtx);

	err := stationsRef.Set(ctx, dataToSet);
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

	fmt.Println("Skonczylem bus: ", busStation);
	cancelFunc();
}

func main(){
	client, err := connect();
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background();

	for i := 1; i <= 26; i++ {
		updateBusStation(ctx, i, client);
	}

}