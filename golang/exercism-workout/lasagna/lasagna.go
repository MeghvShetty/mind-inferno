package lasagna

// TODO: define the 'OvenTime' constant

const (
	OvenTime = 40
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.

func RemainingOvenTime(actualMinutesInOven int) int {

	return OvenTime - actualMinutesInOven

}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.

func PreparationTime(numberOfLayers int) int {

	var b = numberOfLayers * 2

	return b

}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {

	var c = actualMinutesInOven + PreparationTime(numberOfLayers)

	return c

}
