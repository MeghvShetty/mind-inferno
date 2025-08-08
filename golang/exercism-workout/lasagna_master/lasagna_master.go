package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, AvgPreTime int) int {
	if AvgPreTime == 0 {
		AvgPreTime := 2
		return len(layers) * AvgPreTime
	} else {
		return len(layers) * AvgPreTime
	}

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauceLayer int
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles++
		} else if layers[i] == "sauce" {
			sauceLayer++
		}

	}

	return noodles * 50, float64(sauceLayer) * 0.2

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionSize int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		newQuantities[i] = quantities[i] * float64(portionSize) / 2
	}
	return newQuantities
}
