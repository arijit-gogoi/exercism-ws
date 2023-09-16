package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	numerator := int(CalculateWorkingCarsPerHour(productionRate, successRate))
	return int(numerator / 60)
}

// oneCarCost is the cost of a single car.
const oneCarCost = 10000

// tenCarsCost is the cost of ten cars.
const tenCarsCost = 95000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	ones := carsCount % 10
	return uint(tens*tenCarsCost + ones*oneCarCost)
}
