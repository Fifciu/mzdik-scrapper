package main
import "fmt"

func main(){
	firstWay, oppoWay := getBusStations("1");
	for _, element := range firstWay {
		fmt.Println(element.name, element.averageDelay);
	}

	for _, element := range oppoWay {
		fmt.Println(element.name, element.averageDelay);
	}
}