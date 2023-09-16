// Package weather is about the weather.
package weather

var (
	// CurrentCondition is the current condition of the city.
	CurrentCondition string
	// CurrentLocation is the curent location of the city.
	CurrentLocation string
)

// Forecast forecasts the condition in the current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
