// main package of the source code the starting point of the code
package cards

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	FavCards := []int{2, 6, 9}
	return FavCards
}

func OutOfBound(slice []int, index int) bool {

	if index < 0 || index >= len(slice) {

		return true

	} else {

		return false

	}

}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) string {

	if OutOfBound(slice, index) {
		return fmt.Sprint("card == -1")
	} else {
		value := slice[index]
		var vaulue_str string = fmt.Sprint("card == ", value)
		return vaulue_str
	}

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if OutOfBound(slice, index) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	} else {
		return append(values, slice...)
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if OutOfBound(slice, index) {
		fmt.Println("This one")
		return slice
	} else {
		slice := append(slice[:index], slice[index+1:]...)
		return slice
	}
}
