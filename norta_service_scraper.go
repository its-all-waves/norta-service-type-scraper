package norta_service_type_scraper

/*
Scrape https://www.norta.com/rider-tools for service type (vehicle types).

Types
* bus
* streetcar
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

func nortaGetBusRoutes() map[string]string {
	busColl := colly.NewCollector()
	busColl.UserAgent = userAgent

	busColl.OnError(func(r *colly.Response, err error) {
		log.Printf("Bus Route Error: %s", err)
	})
	busColl.OnResponse(func(r *colly.Response) {
		log.Println("--- Buses response received. ---")
	})

	var busRoutesComplete []map[string]string
	busRoutes := make(map[string]string)

	const busRoutesCssSelector = "#hydBusRoutes"
	busColl.OnHTML(busRoutesCssSelector, func(e *colly.HTMLElement) {
		jsonStr := e.Attr("value")
		json.Unmarshal([]byte(jsonStr), &busRoutesComplete)

		for _, info := range busRoutesComplete {
			route := info["RouteCode"]
			name := info["RouteName"]
			busRoutes[route] = name
		}

		// log.Println(busRoutes)
	})

	busColl.Visit("https://www.norta.com/rider-tools")

	return busRoutes
}

func nortaGetStcarRoutes() map[string]string {
	stcarColl := colly.NewCollector()
	stcarColl.UserAgent = userAgent
	stcarColl.OnError(func(r *colly.Response, err error) {
		log.Printf("Streetcar Route Error: %s", err)
	})
	stcarColl.OnResponse(func(r *colly.Response) {
		log.Println("--- Streetcars response received. ---")
	})

	var stcarRoutesComplete []map[string]string
	stcarRoutes := make(map[string]string)

	const stcarRoutesCssSelector = "#hydStreetcarRoutes"
	stcarColl.OnHTML(stcarRoutesCssSelector, func(e *colly.HTMLElement) {
		jsonStr := e.Attr("value")
		json.Unmarshal([]byte(jsonStr), &stcarRoutesComplete)

		for _, info := range stcarRoutesComplete {
			route := info["RouteCode"]
			name := info["RouteName"]
			stcarRoutes[route] = name
		}

		// log.Println(stcarRoutes)
	})

	stcarColl.Visit("https://www.norta.com/rider-tools")

	return stcarRoutes
}

func Scrape() []byte {
	buses := nortaGetBusRoutes()
	stCars := nortaGetStcarRoutes()

	/*
		// starting point
		buses = {
			"103": "103 General Meyer Local",
			"105": "105 Algiers Local"
		}
		streetcars = {
			"12": "12 St. Charles Streetcar",
			"46": "46 Rampart-Loyola Streetcar"
		}
		// goal
		{
		"3": { "type": "bus", "name": "3 Tulane - Elmwood" },
		"8": { "type": "bus", "name": "8 St. Claude - Arabi" },
		"12": { "type": "streetcar", "name": "12 St. Charles Streetcar" },
		"46": { "type": "streetcar", "name": "46 Rampart-Loyola Streetcar" },
		}
	*/
	vehicleTypes := make(map[string]map[string]string)

	for rt, name := range buses {
		vehicleTypes[rt] = map[string]string{
			"type": "bus",
			"name": name,
		}
	}
	for rt, name := range stCars {
		vehicleTypes[rt] = map[string]string{
			"type": "streetcar",
			"name": name,
		}
	}

	vehicleTypesJson, err := json.Marshal(vehicleTypes)
	if err != nil {
		log.Panicln("ERROR: Could not marshal the map into JSON.")
	}

	return vehicleTypesJson
}

func WriteJsonFile() error {
	jsonBytes := Scrape()
	err := os.WriteFile("./output/vehicle_types.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println("ERROR: Could not write the json file to disk.")
		return err
	}
	return nil
}
