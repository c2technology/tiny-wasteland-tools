package character

import "github.com/c2technology/tiny-wasteland-tools/utils"

//Human type Character
const Human = "Human"

//Animal type Character
const Animal = "Animal"

//Leader type Character
const Leader = "Leader"

//WarMachine type Character
const WarMachine = "War Machine"

//things without combat skills
var fodder = []string{
	Animal,
	Human,
	Animal,
	Human,
	Animal,
	Animal,
}

//wild animal or average criminal
var low = []string{
	Human,
	Human,
	Animal,
	Human,
	Animal,
	Human,
}

//dangerous in small groups - skilled combatants
var medium = []string{
	Animal,
	Human,
	Human,
	Human,
	Human,
	Leader,
}

//as dangerous as PCs. Leader of Low or Fodder threats, unique abilities
var high = []string{
	Human,
	Leader,
	Leader,
	Human,
	Leader,
	Human,
}

//more skilled than PCs, a few unique abilities, several Fodder threats
var heroic = []string{
	Leader,
	Leader,
	Leader,
	Leader,
	Leader,
	Human,
}

//avenger threat. highly skilled overpowering to PCs. Can destroy cities
var solo = []string{
	WarMachine,
	Leader,
	WarMachine,
	Leader,
	WarMachine,
	Leader,
}

//RollType for given Character
func RollType(c *Character) {
	switch c.Threat.Name {
	case Fodder.Name:
		c.Type = fodder[utils.Pick(fodder)]
		break
	case Low.Name:
		c.Type = low[utils.Pick(low)]
		break
	case Medium.Name:
		c.Type = medium[utils.Pick(medium)]
		break
	case High.Name:
		c.Type = high[utils.Pick(high)]
		break
	case Heroic.Name:
		c.Type = heroic[utils.Pick(heroic)]
		break
	case Solo.Name:
		c.Type = solo[utils.Pick(solo)]
		break
	}
}
