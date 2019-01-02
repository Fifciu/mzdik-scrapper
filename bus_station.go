package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

type BusStation struct {
	name string
	averageDelay int64
}

func getBusStations(busNumber string) ([]BusStation, []BusStation){
	url := fmt.Sprintf("http://www.mzdik.radom.pl/rozklady/000%s/w.htm", busNumber);
	response, err := http.Get(url);
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close();
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Błąd goquery 1", err)
	}

	// Bus stations' names in first way
	var firstWayStations []BusStation;
	
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(3) > a").Each(func(index int, element *goquery.Selection){
		firstWayStations = append(firstWayStations, BusStation{ 
			name: element.Text(),
			averageDelay: 0,
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(3) > b").Each(func(index int, element *goquery.Selection){
		firstWayStations = append(firstWayStations, BusStation{ 
			name: element.Text(),
			averageDelay: 0,
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(2) > b").Each(func(index int, element *goquery.Selection){
		substr := element.Text()[0:2];
		firstWayStations[index].averageDelay, err = strconv.ParseInt(substr, 10, 8);
		if err != nil {
			log.Fatal("Couldn't parse average delay!")
		}
	})

	// Bus stations' names in opposite way
	var oppoWayStations []BusStation;
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(3) > a").Each(func(index int, element *goquery.Selection){
		oppoWayStations = append(oppoWayStations, BusStation{ 
			name: element.Text(),
			averageDelay: 0,
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(3) > b").Each(func(index int, element *goquery.Selection){
		oppoWayStations = append(oppoWayStations, BusStation{ 
			name: element.Text(),
			averageDelay: 0,
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(2) > b").Each(func(index int, element *goquery.Selection){
		substr := element.Text()[0:2];
		oppoWayStations[index].averageDelay, err = strconv.ParseInt(substr, 10, 8);
		if err != nil {
			log.Fatal("Couldn't parse average delay!")
		}
	})

	return firstWayStations, oppoWayStations;
}