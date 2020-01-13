package character

import (
	"fmt"
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//Human type Character
const Human = "Human"

//Animal type Character
const Animal = "Animal"

//Leader type Character
const Leader = "Leader"

//WarMachine type Character
const WarMachine = "War Machine"

const unknownType = ""

var types = []string{
	Human,
	Animal,
	Leader,
	WarMachine,
}

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

//GetType for given value. If none can be determined an empty type is returned
func GetType(t string) string {
	fmt.Println(fmt.Sprintf("Calculating type for %s", t))
	for _, v := range types {
		if strings.ToLower(v) == strings.ToLower(t) {
			return v
		}
	}
	return unknownType
}

//RollType for given Character if one is not already present.
func RollType(c *Character) {
	fmt.Println(fmt.Sprintf("Type is %s", c.Type))
	if len(c.Type) < 1 {
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
	fmt.Println(fmt.Sprintf("Type is %s", c.Type))
}
