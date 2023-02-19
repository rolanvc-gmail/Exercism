package cars

import (
	"math"
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.00)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var r = math.Floor(float64(productionRate/60) * float64(successRate/100.00))
	return int(r)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var often = int(math.Floor(float64(carsCount) / 10.00))
	var rem = carsCount - 10*often

	return uint(often*95000 + rem*10000)

}
