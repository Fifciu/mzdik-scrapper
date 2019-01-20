package main

import (
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)
 
func main(){
	router := mux.NewRouter();

	router.HandleFunc("/buses", apiGetBuses).Methods("GET");
	router.HandleFunc("/buses/{id}/stations", apiGetBusStations).Methods("GET");
	router.HandleFunc("/buses/{id}/timetable/{way}", apiGetBusTimetable).Methods("GET");
	router.HandleFunc("/buses/{id}/timetable/{way}/{dayType}", apiGetBusDailyTimetable).Methods("GET");

	corsObj := handlers.AllowedOrigins([]string{"*"});

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router))); 
	
	// Updater
	// client, err := dbConnect();
	// if err != nil {
	// 	return;
	// }
	// for i := 1; i <= 26; i++ {
	// 	dbUpdateBusStation(i, client);
	// 	dbUpdateBusTimetable(i, client);
	// }
}