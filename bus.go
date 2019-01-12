package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*type BusStation struct {
	name string
	averageDelay int64
}*/

type BusStation struct {
	Name string
	AverageDelay string
}

type BusTimetable struct {
	CasualDay []BusCourse
	Saturday []BusCourse
	Saints []BusCourse
}

type BusCourse struct {
	Time string
	SpecialLetter rune
}

type BusCourseLetter struct {
	Letter rune
	Description string
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

func getBusTimetable(busNumber string, amount int) (BusTimetable, BusTimetable){
	converted, err := strconv.ParseInt(busNumber, 10, 64);
	if err != nil {
		log.Fatal(err)
	}
	if converted < 10{
		busNumber = "0" + busNumber;
	}

	url := fmt.Sprintf("http://www.mzdik.radom.pl/rozklady/00%s/00%st001.htm", busNumber, busNumber);
	response, err := http.Get(url);
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close();
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Błąd goquery 1", err)
	}

	var Timetable BusTimetable;
	currentDayType := 1;
	var currentHour int64 = 0;
	document.Find("body > font > table > tbody > tr").Each(func(index int, element *goquery.Selection){
		align, _ := element.Attr("align");
		if align == "CENTER" {
			element.Find("td").Each(func(tdIndex int, td *goquery.Selection){
				
				parts := strings.Split(strings.Trim(td.Text(), " "), ".");
				requestedHour, err := strconv.ParseInt(parts[0], 10, 8);
				if err != nil {
					log.Fatal("Strconv err ", err)
				}
				if requestedHour < currentHour {
					currentDayType++;
				}
			
				currentHour = requestedHour;
				letter := td.Text()[len(td.Text()) - 1:];
				_, err = strconv.ParseInt(letter, 10, 8);

				var hour string;

				if err == nil {
					letter = "0"
					hour = strings.Trim(td.Text(), " ");
				} else {
					hour = strings.Trim(td.Text()[0:len(td.Text()) - 1], " ")
				}
				finalLetter := rune(letter[0]);
				if finalLetter == 48 { // "0"
					finalLetter = 0;
				}

		
				switch(currentDayType){
					case 1:
						Timetable.CasualDay = append(Timetable.CasualDay, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
					case 2:
						Timetable.Saturday = append(Timetable.Saturday, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
					case 3:
						Timetable.Saints = append(Timetable.Saints, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
				}
			})
		}
	});

	parsed, err := strconv.ParseInt(busNumber, 10, 64);
	busNumberOffset := strconv.Itoa(int(parsed) + amount - 1);
	url = fmt.Sprintf("http://www.mzdik.radom.pl/rozklady/00%s/00%st0%s.htm", busNumber, busNumber, busNumberOffset);
	response2, err := http.Get(url);
	if err != nil {
		log.Fatal(err)
	}

	defer response2.Body.Close();
	document2, err := goquery.NewDocumentFromReader(response2.Body)
	if err != nil {
		log.Fatal("Błąd goquery 1", err)
	}

	var TimetableBackwards BusTimetable;
	currentDayType = 1;
	currentHour = 0;
	document2.Find("body > font > table > tbody > tr").Each(func(index int, element *goquery.Selection){
		align, _ := element.Attr("align");
		if align == "CENTER" {
			element.Find("td").Each(func(tdIndex int, td *goquery.Selection){
				
				parts := strings.Split(strings.Trim(td.Text(), " "), ".");
				requestedHour, err := strconv.ParseInt(parts[0], 10, 8);
				if err != nil {
					log.Fatal("Strconv err ", err)
				}
				if requestedHour < currentHour {
					currentDayType++;
				}
			
				currentHour = requestedHour;
				letter := td.Text()[len(td.Text()) - 1:];
				_, err = strconv.ParseInt(letter, 10, 8);

				var hour string;

				if err == nil {
					letter = "0"
					hour = strings.Trim(td.Text(), " ");
				} else {
					hour = strings.Trim(td.Text()[0:len(td.Text()) - 1], " ")
				}
				finalLetter := rune(letter[0]);
				if finalLetter == 48 { // "0"
					finalLetter = 0;
				}

		
				switch(currentDayType){
					case 1:
						TimetableBackwards.CasualDay = append(TimetableBackwards.CasualDay, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
					case 2:
						TimetableBackwards.Saturday = append(TimetableBackwards.Saturday, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
					case 3:
						TimetableBackwards.Saints = append(TimetableBackwards.Saints, BusCourse{
							Time: hour,
							SpecialLetter: finalLetter,
						});
				}
			})
		}
	});

	return Timetable, TimetableBackwards;
}