// Package weather provides an i3bar module that displays weather info. 
package weather

// CurrentCondition represents the current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current localisation. 
var CurrentLocation string
// Forecast function represents a string with the CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
