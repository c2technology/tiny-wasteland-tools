package vehicle

import (
	"math/rand"
	"time"
)

var noop = func(*Vehicle) {}

//Upgrade for a Vehicle
type Upgrade struct {
	Name        string
	Description string
	manipulate  vehicleManipulator
}

var vehicleUpgrades = []Upgrade{
	{"Balista", "A giant crossbow mounted on the car that deals 3 damage to its target on success.", noop},
	{"Catapult", "A catapult mounted to the vehicle that deals 1 damage to the enemy driver, all passengers, and the Vehicle on success.", noop},
	{"Emergency Parachute", "A parachute that deploys from the rear of the vehicle to rapidly slow the vehicle. Grants Advantage on all driving Tests that avoid obstacles. After use, all driving Tests gain Disadvantage until the parachute is packed up or removed.", noop},
	{"Flamethrower", "A weapon that sprays a massive gout of flame. Deals 1 damage and the enemy Vehicle and all occupants must make a Save. On failure, they take damage at the start of their next turn and must Save again.", noop},
	{"Grenade Launcher", "A grenade launcher mounted to the vehicle. Deals 4 dam,age to the enemy Vehicle and 1 damage to each occupant.", noop},
	{"Heavy Armor", "Thick metal plating, designed to protect occupants. Reduce all damage dealth to the Vechicle and all occupants by 3 (to a minimum of 1).", noop},
	{"Light Armor", "Light metal plating, providing armor and cover to all occupants. Reduce all damage dealt to the vehicle and occupants by 1 (to a minimum of 1).", noop},
	{"Machine Gun", "A mounted machine gun. Make 3 attacks with Disadvantage on each Action when using this Upgrade.", noop},
	{"Medium Armor", "Plates of metal that protect the vehicle and its occupants. Reduce damage dealth to the vehicle and all occupants by 2 (to a minimum of 1).", noop},
	{"Off-Road Capable", "Improved and upgraded handling makes it easier to drive on bad terrain. You ignore Disadvantage from Rough Terrain while making driving Tests.", noop},
	{"Oil Slick Spray", "Oil that sprays from the back of your car. Any Vehicle following you suffers Disadvantage on their next turn.", noop},
	{"Prisoner Box", "The back of the Vehicle is set up with glass wire and bars. Anyone held in the Prisoner Box in the Vehicle cannot physically interaact with the driver and suffers Disadvantage on any attempts to escape the vehicle.", noop},
	{"Ram Plate", "A massive plate is bolted to the front of your vehicle to increase its impact. Your rams deal 1d6 damage.", noop},
	{"Retrofitted Chassis", "Your chassis has been upgraded with metal and scrap to make it stronger. You vehicle gains +2 HP.",
		func(vehicle *Vehicle) {
			vehicle.HitPoints = vehicle.HitPoints + 2
		},
	}, {"Rocket Launcher", "A launcher tube is attached to your vehicle, allowing you to fire missiles and rockets at others. You may make an Attack that deals 2d3 damage", noop},
	{"Rotary Autocannon", "A massive rotating machine, bolted to your ride. Make 6 Attacks with Disadvantage on each Action. This weapon can only be used once per turn.", noop},
	{"Shredder", "You can eject spikes that shred the tires of those following you. Any Vehicle following you takes 1 damage and has Disadvantage until they stop to repair their tires.", noop},
	{"SMG", "A small machine gun attached to your vehicle. Make 2 Attacks with Disadvantage on each Action. For each attack that hits, you may make an additional Attack with Disadvantage.", noop},
	{"Smoke Dispenser", "Your Vehicle can spray smoke, making escape easier. You can use this system to grant Advantage on rolls to lose someone following you.", noop},
	{"Wheel Blades", "You may make normal melee attacks with your vehicle that are not Ram.", noop},
}

var seed = rand.NewSource(time.Now().UnixNano())
var rando = rand.New(seed)

// SetUpgrades to a Vehicle.
func SetUpgrades(vehicle *Vehicle) {
	for len(vehicle.Upgrades)  < vehicle.maxUpgrades {
		upgrade := vehicleUpgrades[rando.Intn(len(vehicleUpgrades))]
		vehicle.Upgrades[upgrade.Name] = upgrade
		upgrade.manipulate(vehicle)
	}
}
