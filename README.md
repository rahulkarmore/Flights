# Flights
Flights server is used to searching the flights based on "Origin" and "Destination"

## Clone the Project
Clone the project from given link.
https://github.com/rahulkarmore/Flights.git

## Execution Steps

Run the following Command from working dirctory. 
```cmd
go run main.go
```

## Search Flight API (Origin to Destination)
The API is ristricted we need to provied token.

Select POST method and enter given url "http://localhost:8081/r/getFlights".

Please use the following token for authorization.
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlJhaHVsIiwicGFzc3dvcmQiOiJSQGh1bEsiLCJ0b2tlbiI6IiIsImV4cCI6MTYyNDk1NzM5NH0.fA10RzTginH-GPRAam-d1JeUPbCxUfZfigIy0yJJA9M".

Note: To check unauthorised access use differnt token.

## Input
```cmd
{
"origin":"Mumbai",
"destination":"Bhopal"
}
```

# Test Function

Run the following Command from working dirctory. 
```cmd
go test
```



