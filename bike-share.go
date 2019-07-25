package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the application...")

	response, err = http.Get("https://data.melbourne.vic.gov.au/resource/tdvh-n9dv.json")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)

		var bd []BikeData

		err := json.Unmarshal(data, &bd)
		if err != nil {
			log.Fatal(err)
		}

		for index, element := range bd {
			fmt.Println(index, "Station Id:", element.StationID, "\tAvailable Bikes: ", element.AvailableBikes, "\t Capacity: ", element.Capacity)
		}

	}
	fmt.Println("Terminating the application...")
}

type BikeData struct {
	StationID      string `json:"station_id"`
	AvailableBikes string `json:"available_bikes"`
	EmptyDocks     string `json:"empty_docks"`
	Capacity       string `json:"capacity"`
	LastUpdated    string `json:"last_updated"`
}
