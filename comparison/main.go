package purchase
import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car") || (kind == "truck") {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    str := " is clearly the better choice."
    if strings.Compare(option1,option2) < 0 {
        return option1 + str
    }
    return option2 + str
	
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
        return originalPrice * 0.80
    } else if age < 10 {
        return originalPrice * 0.70
    }
    return originalPrice * 0.50
}
