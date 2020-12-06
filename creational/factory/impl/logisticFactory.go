package factory

import "fmt"

// Logistic return the right instance
func Logistic(transportType string) (Logistics, error) {
	if transportType == "Truck" {
		return NewTruck(), nil
	}

	if transportType == "Boat" {
		return NewBoat(), nil
	}

	return nil, fmt.Errorf("Wrong transport type")
}
