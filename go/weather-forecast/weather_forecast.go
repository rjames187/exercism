// Package weather contains one function for generating a forecast.
package weather

// CurrentCondition stores the current condition of the weather; can be used for forecasts.
var CurrentCondition string
// CurrentLocation stores the current location of the weather, can be used for forecasts.
var CurrentLocation string

// Forecast generates a weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
