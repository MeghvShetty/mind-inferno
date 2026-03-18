package lasagna

// TODO: define the 'OvenTime' constant

var OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remainingTime := OvenTime - actualMinutesInOven
	return remainingTime 
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	preparation := numberOfLayers * 2
	return preparation
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven	
	return elapsedTime 
}
