package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func apiGetBuses(w http.ResponseWriter, r *http.Request){
	buses := dbGetBuses();
	json.NewEncoder(w).Encode(buses);
}

func apiGetBusStations(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, err := strconv.ParseInt(params["id"], 10, 64);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	client, err := dbConnect();
	if err != nil {
		http.Error(w, "Couldn't connect to Database, sorry", http.StatusBadRequest);
		return;
	}

	forward, backward := dbGetBusStations(int(id), client);
	returnable := make(map[string][]BusStation)
	returnable["Forward"] = forward;
	returnable["Backward"] = backward;
	json.NewEncoder(w).Encode(returnable);
}

func apiGetBusTimetable(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, err := strconv.ParseInt(params["id"], 10, 64);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	if params["way"] != "forward" && params["way"] != "backward" {
		http.Error(w, "Unproper way", http.StatusBadRequest);
		return;
	}

	backwards := false
	if params["way"] == "backward" {
		backwards = true;
	}

	client, err := dbConnect();
	if err != nil {
		http.Error(w, "Couldn't connect to Database, sorry", http.StatusBadRequest);
		return;
	}

	timetable := dbGetBusFullTimetable(int(id), backwards, client);
	json.NewEncoder(w).Encode(timetable);
}

func apiGetBusDailyTimetable(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, err := strconv.ParseInt(params["id"], 10, 64);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	if params["way"] != "forward" && params["way"] != "backward" {
		http.Error(w, "Unproper way", http.StatusBadRequest);
		return;
	}

	backwards := false
	if params["way"] == "backward" {
		backwards = true;
	}

	dayType, err := strconv.ParseInt(params["dayType"], 10, 64);
	if err != nil || dayType < 1 || dayType > 3{
		http.Error(w, "Unproper daytype", http.StatusBadRequest);
		return;
	}

	client, err := dbConnect();
	if err != nil {
		http.Error(w, "Couldn't connect to Database, sorry", http.StatusBadRequest);
		return;
	}

	timetable := dbGetBusCertainTimetable(int(id), backwards, client, dayType);
	json.NewEncoder(w).Encode(timetable); 
}