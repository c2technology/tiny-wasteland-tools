package vehicle

//vehicleManipulator that manipulates various attributes of a vehicle
type vehicleManipulator func(*Vehicle)

//Chassis for a vehicle
type Chassis struct {
	Name        string
	Description string
	manipulate  vehicleManipulator
}

//SetChassis for the given Vehicle
func (chassis Chassis) set(vehicle *Vehicle) {
	vehicle.Chassis = chassis
	chassis.manipulate(vehicle)
}

var vehicleChassis = []Chassis{
	{"Motorcycle", "", func(vehicle *Vehicle) {
		vehicle.HitPoints = 4
		vehicle.Capacity = vehicle.Capacity + 1
		vehicle.maxUpgrades = vehicle.maxUpgrades + 1
		vehicle.Upgrades["Evasive"] = Upgrade{Name: "Evasive", Description: "Whenever you are attacked, you may Test. On success, the attack misses. This does not stack with Evade."}
	}},
	{"Sedan", "", func(vehicle *Vehicle) {
		vehicle.HitPoints = 6
		vehicle.Capacity = vehicle.Capacity + 4
		vehicle.maxUpgrades = vehicle.maxUpgrades + 1
	}},
	{"Muscle Car", "", func(vehicle *Vehicle) {
		vehicle.HitPoints = 8
		vehicle.Capacity = vehicle.Capacity + 3
		vehicle.maxUpgrades = vehicle.maxUpgrades + 1
		vehicle.Upgrades["Detroit Steel"] = Upgrade{Name: "Detroit Steel", Description: "You take no damage from Rams you initiate."}
	}},
	{"Truck", "", func(vehicle *Vehicle) {
		vehicle.HitPoints = 8
		vehicle.Capacity = vehicle.Capacity + 5
		vehicle.maxUpgrades = vehicle.maxUpgrades + 1
		vehicle.Upgrades["4-Wheel Drive"] = Upgrade{Name: "4-Wheel Drive", Description: "You ignore Disadvantage on driving Tests from Rough Terrain"}
	}},
	{"BFV", "", func(vehicle *Vehicle) {
		vehicle.HitPoints = 14
		vehicle.Capacity = vehicle.Capacity + 12
		vehicle.maxUpgrades = 0
	}},
}

//SetChassis randomly selected from the default, for a vehicle
func SetChassis(vehicle *Vehicle) {
	chassis := vehicleChassis[rando.Intn(len(vehicleChassis))]
	vehicle.Chassis = chassis
	chassis.manipulate(vehicle)
}
