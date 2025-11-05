package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentage := successRate/100
    return float64(productionRate) *  percentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percent := successRate/100
    perHour:= float64(productionRate) *  percent
    return int(perHour/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    byTens := carsCount/10
    remainder := carsCount % 10
	totalCost := byTens*95000 + remainder*10000
    return uint(totalCost)
}
