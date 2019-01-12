package main

import (
	//"context"
	"log"
	/*"firebase.google.com/go/db"
	"strconv"*/
	//"fmt"
)

func main(){
	client, err := connect();
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Updating bus stations list 1-26
	// Updating timetable both side
	for i := 1; i <= 26; i++ {
		updateBusStation(i, client);
		updateBusTimetable(i, client);
	}

	//test := getBusTimetable("1");
	//fmt.Println(test);
	//updateBusTimetable(1, client)
}