package model

type Flights struct {
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Date        string    `json:"date"`
	DepTime     string    `json:"deptTime"`
	ArrTime     string    `json:"arrTime"`
	FlightNo    string    `json:"flightNo"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	LayOver     []Flights `json:"layOver"`
}

type FlightsList []Flights
