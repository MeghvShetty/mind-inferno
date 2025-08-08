package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists_unit := units[unit]
	if !exists_unit {
		return false
	}
	_, exists_item := bill[item]
	if !exists_item {
		return false
	}
	if bill[item] < units[unit] {
		return false
	}
	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
		return true
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
