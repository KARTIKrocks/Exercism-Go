// Package weather which contains all the weather related variables and functitions.
package weather

// CurrentCondition shows the current condition of Goblinocus.
var CurrentCondition string

// CurrentLocation shows shows current location of the city of the Goblinocus.
var CurrentLocation string

// Forecast shows the forecast of the particular city of Goblinocus according to its condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
