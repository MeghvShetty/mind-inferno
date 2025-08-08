package gross

import "testing"

func TestUnits(t *testing.T) {
	units := Units()
	expected := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	for k, v := range expected {
		if units[k] != v {
			t.Errorf("Units()[%s] = %d; want %d", k, units[k], v)
		}
	}
}

func TestNewBill(t *testing.T) {
	bill := NewBill()
	if bill == nil {
		t.Error("NewBill() returned nil map")
	}
	if len(bill) != 0 {
		t.Error("NewBill() returned non-empty map")
	}
}

func TestAddItem(t *testing.T) {
	bill := NewBill()
	units := Units()

	// Test adding valid item
	success := AddItem(bill, units, "carrot", "dozen")
	if !success {
		t.Error("AddItem() failed to add valid item")
	}
	if bill["carrot"] != 12 {
		t.Errorf("AddItem() added wrong quantity: got %d, want 12", bill["carrot"])
	}

	// Test adding invalid unit
	success = AddItem(bill, units, "carrot", "invalid_unit")
	if success {
		t.Error("AddItem() succeeded with invalid unit")
	}
}

func TestRemoveItem(t *testing.T) {
	bill := NewBill()
	units := Units()

	// Add items first
	AddItem(bill, units, "carrot", "dozen")
	AddItem(bill, units, "apple", "half_of_a_dozen")

	// Test removing valid item
	success := RemoveItem(bill, units, "carrot", "quarter_of_a_dozen")
	if !success {
		t.Error("RemoveItem() failed to remove valid item")
	}
	if bill["carrot"] != 9 {
		t.Errorf("RemoveItem() removed wrong quantity: got %d, want 9", bill["carrot"])
	}

	// Test removing non-existent item
	success = RemoveItem(bill, units, "banana", "dozen")
	if success {
		t.Error("RemoveItem() succeeded with non-existent item")
	}

	// Test removing with invalid unit
	success = RemoveItem(bill, units, "carrot", "invalid_unit")
	if success {
		t.Error("RemoveItem() succeeded with invalid unit")
	}

	// Test removing more than available
	success = RemoveItem(bill, units, "carrot", "dozen")
	if success {
		t.Error("RemoveItem() succeeded when removing more than available")
	}
}

func TestGetItem(t *testing.T) {
	bill := NewBill()
	units := Units()

	// Add an item
	AddItem(bill, units, "carrot", "dozen")

	// Test getting existing item
	quantity, exists := GetItem(bill, "carrot")
	if !exists {
		t.Error("GetItem() failed to find existing item")
	}
	if quantity != 12 {
		t.Errorf("GetItem() returned wrong quantity: got %d, want 12", quantity)
	}

	// Test getting non-existent item
	quantity, exists = GetItem(bill, "banana")
	if exists {
		t.Error("GetItem() found non-existent item")
	}
	if quantity != 0 {
		t.Errorf("GetItem() returned non-zero quantity for non-existent item: got %d", quantity)
	}
}
