package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter();

	router.HandleFunc("/buses", apiGetBuses).Methods("GET");
	router.HandleFunc("/buses/{id}/stations", apiGetBusStations).Methods("GET");
	router.HandleFunc("/buses/{id}/timetable/{way}", apiGetBusTimetable).Methods("GET");
	router.HandleFunc("/buses/{id}/timetable/{way}/{dayType}", apiGetBusDailyTimetable).Methods("GET");

	log.Fatal(http.ListenAndServe(":8000", router));
	/*client, err := dbConnect();
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	for i := 1; i <= 26; i++ {
		dbUpdateBusStation(i, client);
		dbUpdateBusTimetable(i, client);
	}*/
}