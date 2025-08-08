package cars

// CalculateWorkingCarsPerHour calculates how many working cars are

// produced by the assembly line every hour.

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	return float64(productionRate) * successRate / 100

}

// CalculateWorkingCarsPerMinute calculates how many working cars are

// produced by the assembly line every minute.

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	var value_3 = CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(value_3) / 60

}

const b uint = 10

const produce_total = 95000

const per_cost = 10000

// CalculateCost works out the cost of producing the given number of cars.

func CalculateCost(carsCount int) uint {

	var value_1 = uint(carsCount) % b

	var value_2 int = (carsCount) / 10

	var cost_ten uint = uint(value_2) * produce_total

	var cost_per uint = value_1 * per_cost

	return cost_ten + cost_per

}
