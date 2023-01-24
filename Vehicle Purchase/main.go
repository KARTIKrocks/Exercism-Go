package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	result := strings.Compare(option1, option2)
	endString := " is clearly the better choice."
	if result == -1 {
		return option1 + endString
	} else {
		return option2 + endString
	}

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
		result := originalPrice * 0.5
		return result
	} else if age >= 3 {
		result := originalPrice * 0.7
		return result
	} else {
		// it is for age < 3
		result := originalPrice * 0.8
		return result
	}
}
