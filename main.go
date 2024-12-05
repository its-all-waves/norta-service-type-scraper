package main

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

var htmlElemIds = map[string]string{
	"ferry":     "#hydFerryRoutes",
	"streetcar": "#hydStreetcarRoutes",
	"bus":       "#hydBusRoutes",
}

func nortaGetRoutes(vehType string) map[string]string {
	collector := colly.NewCollector()
	collector.UserAgent = userAgent

	collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Bus Route Error: %s", err)
	})
	collector.OnResponse(func(r *colly.Response) {
		log.Println("--- Buses response received. ---")
	})

	var routesComplete []map[string]string
	routes := make(map[string]string)
	query := htmlElemIds[vehType]
	collector.OnHTML(
		query,
		func(e *colly.HTMLElement) {
			jsonStr := e.Attr("value")
			json.Unmarshal([]byte(jsonStr), &routesComplete)
			for _, info := range routesComplete {
				route := info["RouteCode"]
				name := info["RouteName"]
				routes[route] = name
			}
			// log.Println(routes)
		},
	)

	collector.Visit("https://www.norta.com/rider-tools")

	return routes
}

func Scrape() []byte {
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

	for rt, name := range nortaGetRoutes("bus") {
		vehicleTypes[rt] = map[string]string{
			"type": "bus",
			"name": name,
		}
	}
	for rt, name := range nortaGetRoutes("streetcar") {
		vehicleTypes[rt] = map[string]string{
			"type": "streetcar",
			"name": name,
		}
	}
	for rt, name := range nortaGetRoutes("ferry") {
		vehicleTypes[rt] = map[string]string{
			"type": "ferry",
			"name": name,
		}
	}

	vehicleTypesJson, err := json.Marshal(vehicleTypes)
	if err != nil {
		log.Panicln("ERROR: Could not marshal the map into JSON.")
	}

	return vehicleTypesJson
}

func Write(jsonBytes []byte) error {
	err := os.WriteFile("./output/routes_info.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println("ERROR: Could not write the json file to disk.", err)
		return err
	}
	return nil
}

func main() {
	result := Scrape()
	Write(result)
}
