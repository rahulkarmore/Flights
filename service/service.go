package service

import (
	"Flights/model"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func GetFlightsService(requestData model.Flights) (flight model.FlightsList) {
	fmt.Println("request data---", requestData)

	dummyData := `[
    {"origin":"Mumbai", "destination":"Delhi", "date":"01/06/2021", "depTime":"5:00","arrTime":"8:30","flightNo":"IND-46","name":"Indigo","price":8000},
    {"origin":"Mumbai", "destination":"Banglore", "date":"01/06/2021", "depTime":"6:00","arrTime":"9:30","flightNo":"IND-47","name":"Indigo","price":8000},
    {"origin":"Mumbai", "destination":"Pune", "date":"01/06/2021", "depTime":"7:30","arrTime":"8:45","flightNo":"GO-107","name":"GO Air","price":8000},
    {"origin":"Mumbai", "destination":"Nagpur", "date":"01/06/2021", "depTime":"16:00","arrTime":"18:00","flightNo":"GO-108","name":"Go Air","price":8000},
    
    {"origin":"Pune", "destination":"Delhi", "date":"01/06/2021", "depTime":"13:00","arrTime":"15:00","flightNo":"IND-47","name":"Indigo","price":8000},
    {"origin":"Pune", "destination":"Banglore", "date":"01/06/2021", "depTime":"16:30","arrTime":"19:10","flightNo":"IND-48","name":"Indigo","price":8000},
    {"origin":"Pune", "destination":"Nagpur", "date":"01/06/2021", "depTime":"20:30","arrTime":"21:15","flightNo":"GO-109","name":"GO Air","price":8000},
    {"origin":"Pune", "destination":"Mumbai", "date":"01/06/2021", "depTime":"22:00","arrTime":"23:00","flightNo":"GO-110","name":"Go Air","price":8000},
    {"origin":"Pune", "destination":"Indore", "date":"01/06/2021", "depTime":"21:00","arrTime":"19:00","flightNo":"GO-310","name":"Go Air","price":8000},
    {"origin":"Pune", "destination":"Bhopal", "date":"01/06/2021", "depTime":"21:00","arrTime":"20:00","flightNo":"GO-320","name":"Go Air","price":8000},
    
    {"origin":"Banglore", "destination":"Delhi", "date":"01/06/2021", "depTime":"4:00","arrTime":"6:00","flightNo":"IND-50","name":"Indigo","price":8000},
    {"origin":"Banglore", "destination":"Pune", "date":"01/06/2021", "depTime":"20:30","arrTime":"22:10","flightNo":"IND-51","name":"Indigo","price":8000},
    {"origin":"Banglore", "destination":"Nagpur", "date":"01/06/2021", "depTime":"16:30","arrTime":"22:15","flightNo":"GO-109","name":"GO Air","price":8000},
    {"origin":"Banglore", "destination":"Mumbai", "date":"01/06/2021", "depTime":"10:00","arrTime":"11:30","flightNo":"GO-112","name":"Go Air","price":8000},
    {"origin":"Banglore", "destination":"Mumbai", "date":"01/06/2021", "depTime":"13:00","arrTime":"15:30","flightNo":"GO-114","name":"Go Air","price":8000},
    
    
    {"origin":"Indore", "destination":"Pune", "date":"01/06/2021", "depTime":"6:00","arrTime":"9:40","flightNo":"GO-116","name":"Go Air","price":8000},
    {"origin":"Indore", "destination":"Mumbai", "date":"01/06/2021", "depTime":"13:00","arrTime":"16:00","flightNo":"GO-117","name":"Go Air","price":8000},
    {"origin":"Indore", "destination":"Bhopal", "date":"01/06/2021", "depTime":"22:00","arrTime":"23:40","flightNo":"GO-119","name":"Go Air","price":8000},
    
    {"origin":"Bhopal", "destination":"Pune", "date":"01/06/2021", "depTime":"7:00","arrTime":"9:40","flightNo":"GO-118","name":"Go Air","price":8000},
    {"origin":"Bhopal", "destination":"Mumbai", "date":"01/06/2021", "depTime":"13:00","arrTime":"16:00","flightNo":"GO-120","name":"Go Air","price":8000},
    {"origin":"Bhopal", "destination":"Indore", "date":"01/06/2021", "depTime":"23:00","arrTime":"23:45","flightNo":"GO-122","name":"Go Air","price":8000}
 ]`

	var FlightsData, result model.FlightsList

	err := json.Unmarshal([]byte(dummyData), &FlightsData)
	if err != nil {
		log.Println("error in unmarshal in file", err)

	}
	// fmt.Println(" flight data", FlightsData)
	origin := strings.Trim(requestData.Origin, " ")
	destination := strings.Trim(requestData.Destination, " ")
	for i := 0; i < len(FlightsData); i++ {

		if origin == FlightsData[i].Origin && destination == FlightsData[i].Destination {
			flightDetails := model.Flights{
				Origin:      origin,
				Destination: destination,
				Name:        FlightsData[i].Name,
				FlightNo:    FlightsData[i].FlightNo,
				ArrTime:     FlightsData[i].ArrTime,
				DepTime:     FlightsData[i].DepTime,
				Date:        FlightsData[i].Date,
			}

			result = append(result, flightDetails)
			continue

		} else {
			for j := i + 1; j < len(FlightsData); j++ {

				if origin == FlightsData[i].Origin && FlightsData[i].Destination == FlightsData[j].Origin && FlightsData[j].Destination == destination {

					var layOver []model.Flights
					layOver = append(layOver, FlightsData[i])
					layOver = append(layOver, FlightsData[j])

					flightDetails := model.Flights{
						Origin:      origin,
						Destination: destination,
						Name:        FlightsData[i].Name,
						FlightNo:    FlightsData[i].FlightNo,
						ArrTime:     FlightsData[i].ArrTime,
						DepTime:     FlightsData[i].DepTime,
						Date:        FlightsData[i].Date,
						LayOver:     layOver,
					}

					result = append(result, flightDetails)

				}
			}
		}

	}

	fmt.Println("Result : ", result)

	return result

}
