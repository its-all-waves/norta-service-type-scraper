package norta_service_type_scraper

import "testing"

func TestNortaGetBusRoutes(t *testing.T) {
	busRoutes := nortaGetBusRoutes()
	if len(busRoutes) == 0 {
		t.Fatalf("No bus routes were retrieved.")
	}
}

func TestNortaGetStcarRoutes(t *testing.T) {
	stCarRoutes := nortaGetStcarRoutes()
	if len(stCarRoutes) == 0 {
		t.Fatalf("No streetcar routes were received.")
	}
}
