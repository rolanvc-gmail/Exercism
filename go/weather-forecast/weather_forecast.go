// Package weather provides the Forecast() function which prints the
// weather condtion in the specified city.
package weather

// CurrentCondition is used to describe the condition of the city in the Forecast() function.
var CurrentCondition string

// CurrentLocation is used as the city location for the Forecast function.
var CurrentLocation string

// Forecast() takes the city and the condition as inputs
// and outputs a string describing the weather condition in that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
