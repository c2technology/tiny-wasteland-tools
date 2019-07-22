package main

import (
	"fmt"
	"github.com/c2technology/tiny-wasteland-tools/vehicle"
)

// Generates a vehicle
func main() {
	vehicle := vehicle.Generate()
	//TODO: Pull name from input or a name generator
	fmt.Println(fmt.Sprintf("Chassis: %s", vehicle.Chassis.Name))
	fmt.Println(fmt.Sprintf("Capacity: %d", vehicle.Capacity))
	fmt.Println(fmt.Sprintf("Hit Points: %d", vehicle.HitPoints))
	fmt.Println("Upgrades:")
	for key, val := range vehicle.Upgrades {
		fmt.Println(fmt.Sprintf("  %s: %s", key, val.Description))
	}
}
