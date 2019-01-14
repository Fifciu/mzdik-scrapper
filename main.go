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
}