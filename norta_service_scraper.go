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

	"github.com/gocolly/colly"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

func nortaGetBusRoutes() map[string]string {
	busColl := colly.NewCollector()
	busColl.UserAgent = userAgent

	busColl.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Bus Route Error: %s", err)
	})
	busColl.OnResponse(func(r *colly.Response) {
		fmt.Println("--- Bus routes response received. ---")
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

		fmt.Println(busRoutes)
	})

	busColl.Visit("https://www.norta.com/rider-tools")

	return busRoutes
}

func nortaGetStcarRoutes() map[string]string {
	stcarColl := colly.NewCollector()
	stcarColl.UserAgent = userAgent
	stcarColl.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Streetcar Route Error: %s", err)
	})
	stcarColl.OnResponse(func(r *colly.Response) {
		fmt.Println("--- Streetcar routes response received. ---")
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

		fmt.Println(stcarRoutes)
	})

	stcarColl.Visit("https://www.norta.com/rider-tools")

	return stcarRoutes
}
