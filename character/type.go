package character

import "github.com/c2technology/tiny-wasteland-tools/utils"

const HUMAN = "Human"
const ANIMAL = "Animal"
const LEADER = "Leader"
const WAR_MACHINE = "War Machine"

//things without combat skills
var fodder = []string{
	ANIMAL,
	HUMAN,
	ANIMAL,
	HUMAN,
	ANIMAL,
	ANIMAL,
}

//wild animal or average criminal
var low = []string{
	HUMAN,
	HUMAN,
	ANIMAL,
	HUMAN,
	ANIMAL,
	HUMAN,
}

//dangerous in small groups - skilled combatants
var medium = []string{
	ANIMAL,
	HUMAN,
	HUMAN,
	HUMAN,
	HUMAN,
	LEADER,
}

//as dangerous as PCs. Leader of Low or Fodder threats, unique abilities
var high = []string{
	HUMAN,
	LEADER,
	LEADER,
	HUMAN,
	LEADER,
	HUMAN,
}

//more skilled than PCs, a few unique abilities, several Fodder threats
var heroic = []string{
	LEADER,
	LEADER,
	LEADER,
	LEADER,
	LEADER,
	HUMAN,
}

//avenger level threats. highly skilled overpowering to PCs. Can level cities
var solo = []string{
	WAR_MACHINE,
	LEADER,
	WAR_MACHINE,
	LEADER,
	WAR_MACHINE,
	LEADER,
}

func RollType(c *Character) {
	switch c.Level.Name {
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
