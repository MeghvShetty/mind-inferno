// Package weather forecast weather for Goblinocus.

// It uses the location to give weather report.

package weather

// CurrentCondition assigns the wether status to the variabele.

var CurrentCondition string

// CurrentLocation assigns location for the place.

var CurrentLocation string

//Forecast function retuns the currentloaction and CurrentCondition.

func Forecast(city, condition string) string {

	CurrentLocation, CurrentCondition = city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition

}
