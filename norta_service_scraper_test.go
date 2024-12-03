package norta_vehicle_types

import (
	"log"
	"testing"
)

func TestNortaGetBusRoutes(t *testing.T) {
	busRoutes := nortaGetBusRoutes()
	// log.Println("Bus Routes:", busRoutes)
	if len(busRoutes) == 0 {
		t.Fatalf("No bus routes were retrieved.")
	}
}

func TestNortaGetStcarRoutes(t *testing.T) {
	stCarRoutes := nortaGetStcarRoutes()
	// log.Println("Streetcar Routes:", stCarRoutes)
	if len(stCarRoutes) == 0 {
		t.Fatalf("No streetcar routes were received.")
	}
}

func TestScraper(t *testing.T) {
	result := Scrape()
	log.Println("Result:", result)
	if len(result) == 0 {
		t.Fatalf("No results were produced.")
	}
}
