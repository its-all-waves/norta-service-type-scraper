package main

import (
	"log"
	"testing"
)

func TestNortaBusRoutes(t *testing.T) {
	routes := nortaGetRoutes("bus")
	// log.Println("Bus Routes:", busRoutes)
	if len(routes) == 0 {
		t.Fatalf("No bus routes were retrieved.")
	}
}

func TestNortaStCarRoutes(t *testing.T) {
	routes := nortaGetRoutes("streetcar")
	// log.Println("Streetcar Routes:", stCarRoutes)
	if len(routes) == 0 {
		t.Fatalf("No streetcar routes were received.")
	}
}

func TestNortaFerryRoutes(t *testing.T) {
	routes := nortaGetRoutes("ferry")
	// log.Println("Ferry Routes:", stCarRoutes)
	if len(routes) == 0 {
		t.Fatalf("No ferry routes were received.")
	}
}

func TestScraper(t *testing.T) {
	result := Scrape()
	log.Println("Result:", string(result))
	if len(result) == 0 {
		t.Fatalf("No results were produced.")
	}
}
