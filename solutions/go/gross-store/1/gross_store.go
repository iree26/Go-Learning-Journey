package gross


// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	list := map[string]int{
       "quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
    }
    
  return list
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := map[string]int{}
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := units[unit]
    if !ok {
        return false // the unit is not valid
    }

    // If the item already exists in the bill, increase its quantity
    bill[item] += qty

    return true // success
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := units[unit]
	if !ok {
		return false // invalid unit
	}

	itemQty, ok := bill[item]
	if !ok {
		return false // item not found
	}

	newQty := itemQty - qty
	if newQty < 0 {
		return false // can’t go below zero
	} else if newQty == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newQty
	return true
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    qty, ok := bill[item]
    if !ok {
        return 0, false
    }
    return qty, true
}

