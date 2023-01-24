package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	result := float64(productionRate) * successRate / 100
	return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	result := float64(productionRate) * successRate / (100 * 60)
	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	CostOfGroupOfTenCars := 95000
	CostOfIndividualCar := 10000
	remainder := carsCount % 10
	
	switch carsCount {
	case 0:
		return 0 
	case 1,2,3,4,5,6,7,8,9:
		return uint(carsCount * 10000)
	}

	// we are given in instructin that
	// 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars
	// So the cost for 37 cars is: 3*95,000+7*10,000=355,000
	result := (int(carsCount/10) * CostOfGroupOfTenCars) + (remainder * CostOfIndividualCar)
	return uint(result)
}
