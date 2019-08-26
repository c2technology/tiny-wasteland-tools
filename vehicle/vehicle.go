package vehicle

//Vehicle contains defined attributes
type Vehicle struct {
	HitPoints   int
	Capacity    int
	Chassis     Chassis
	Upgrades    map[string]Upgrade
	maxUpgrades int
}

//Generate a Vehicle with random attributes and the given name
func Generate() Vehicle {
	vehicle := Vehicle{
		Capacity:    1,
		Upgrades:    make(map[string]Upgrade),
		maxUpgrades: 3,
	}
	SetChassis(&vehicle)
	SetUpgrades(&vehicle)
	//	setInventory(&Vehicle)
	//	setClix(&Vehicle)
	//	setProficiency(&Vehicle)
	//	setMastery(&Vehicle)
	return vehicle
}