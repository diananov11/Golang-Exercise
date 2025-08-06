// Package weather for weather forecast in your location.
package weather

// CurrentCondition stores the weather condition.
var CurrentCondition string
// CurrentLocation stores the location of the city.
var CurrentLocation string

// Forecast returns a string value equal to the weather forecast of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
