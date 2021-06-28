package main

import (
	"Flights/model"
	"Flights/service"
	"testing"
)

func TestGetFlights(t *testing.T) {

	flightData := model.Flights{
		Origin:      "Pune",
		Destination: "Nagpur",
	}

	result := service.GetFlightsService(flightData)

	if result == nil {

		t.Error("Result not found :", result)
	}

}
