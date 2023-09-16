package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	a := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return a
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	n, y := units[unit]
	if y == false {
		return false
	}
	bill[item] += n
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billAmount, billOk := bill[item]
	unitAmount, unitOk := units[unit]

	if !billOk || !unitOk || billAmount < unitAmount {
		return false
	}

	if billAmount > unitAmount {
		bill[item] = billAmount - unitAmount
	} else {
		delete(bill, item)
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	q, ok := bill[item]
	return q, ok
}
