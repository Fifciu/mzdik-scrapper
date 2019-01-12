package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

/*type BusStation struct {
	name string
	averageDelay int64
}*/

type BusStation struct {
	Name string
	AverageDelay string
}

func getBusStations(busNumber string) ([]BusStation, []BusStation){
	converted, err := strconv.ParseInt(busNumber, 10, 64);
	if err != nil {
		log.Fatal(err)
	}
	if converted < 10{
		busNumber = "0" + busNumber;
	}

	url := fmt.Sprintf("http://www.mzdik.radom.pl/rozklady/00%s/w.htm", busNumber);
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
			Name: element.Text(),
			AverageDelay: "0",
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(3) > b").Each(func(index int, element *goquery.Selection){
		firstWayStations = append(firstWayStations, BusStation{ 
			Name: element.Text(),
			AverageDelay: "0",
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(1) > table > tbody > tr > td:nth-child(2) > b").Each(func(index int, element *goquery.Selection){
		substr := element.Text()[0:2];
		tmp, err := strconv.ParseInt(substr, 10, 8);
		tmp2 := int(tmp);
		firstWayStations[index].AverageDelay = strconv.Itoa(tmp2);
		if err != nil {
			log.Fatal("Couldn't parse average delay!")
		}
	})

	// Bus stations' names in opposite way
	var oppoWayStations []BusStation;
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr > td:nth-child(3) > a").Each(func(index int, element *goquery.Selection){
		oppoWayStations = append(oppoWayStations, BusStation{ 
			Name: element.Text(),
			AverageDelay: "0",
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr > td:nth-child(3) > b").Each(func(index int, element *goquery.Selection){
		oppoWayStations = append(oppoWayStations, BusStation{ 
			Name: element.Text(),
			AverageDelay: "0",
		});
	})
	document.Find("body > font > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr > td:nth-child(2) > b").Each(func(index int, element *goquery.Selection){
		substr := element.Text()[0:2];
		tmp, err := strconv.ParseInt(substr, 10, 8);
		tmp2 := int(tmp);
		oppoWayStations[index].AverageDelay = strconv.Itoa(tmp2);
		if err != nil {
			log.Fatal("Couldn't parse average delay!")
		}
	})

	return firstWayStations, oppoWayStations;
}