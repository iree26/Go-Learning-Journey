//Package weather contains tools that 
// can forecast the weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition represents the current condition to be recorded.
	CurrentCondition string
    // CurrentLocation represents the current location to be recorded.
	CurrentLocation  string
)

// Forecast returns a string that shows the current location and the current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
