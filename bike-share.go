package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func main() {

	var err error
	var response *http.Response
	var bd []BikeData

	fmt.Println("Starting the application...")

	response, err = http.Get("https://data.melbourne.vic.gov.au/resource/tdvh-n9dv.json")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Convert json to BikeData
		err = json.Unmarshal(data, &bd)
		if err != nil {
			log.Fatal(err)
		}

		sort.Slice(bd[:], func(i, j int) bool {
			return bd[i].StationID < bd[j].StationID
		})

		// Write out table of all bikes in Melbourne
		for index, element := range bd {
			fmt.Println(index, "Station Id:", element.StationID, "\tAvailable Bikes: ", element.AvailableBikes, "\t Capacity: ", element.Capacity)
		}

	}
	fmt.Println("Terminating the application...")
}

type BikeData struct {
	StationID      int64  `json:"station_id,string"`
	AvailableBikes string `json:"available_bikes"`
	EmptyDocks     string `json:"empty_docks"`
	Capacity       string `json:"capacity"`
	LastUpdated    string `json:"last_updated"`
}
