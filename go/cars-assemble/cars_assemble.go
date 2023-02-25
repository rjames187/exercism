package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percent := successRate / 100.0
	result := float64(productionRate) * percent
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	percent := successRate / 100.0
	productionPerMinute := productionRate / 60
	result := float64(productionPerMinute) * percent
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	remainder := carsCount % 10
	numTens := (carsCount - remainder) / 10
	tensCost := numTens * 95000
	remainderCost := remainder * 10000
	return uint(tensCost + remainderCost)
}
