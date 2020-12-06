package main

import (
	"fmt"

	factory "github.com/lordkevinmo/go_design_pattern/creational/factory/impl"
)

func main() {
	boat, err := factory.Logistic("Boat")
	if err != nil {
		fmt.Println(err)
	}
	truck, err := factory.Logistic("Truck")
	if err != nil {
		fmt.Println(err)
	}

	displayDetails(boat)
	displayDetails(truck)
}

func displayDetails(log factory.Logistics) {
	val := fmt.Sprintf("Transport with name %s and engine %d power", log.Name(), log.Engine())
	fmt.Println(val)
}
