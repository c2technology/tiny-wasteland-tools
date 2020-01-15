package character

import (
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

const human = "Human"

const animal = "Animal"

const leader = "Leader"

const warMachine = "War Machine"

const unknownRace = ""

var characterTypes = []string{
	human,
	animal,
	leader,
	warMachine,
}

//things without combat skills
var fodderCharacterTypes = []string{
	animal,
	human,
	animal,
	human,
	animal,
	animal,
}

//wild animal or average criminal
var lowCharacterTypes = []string{
	human,
	human,
	animal,
	human,
	animal,
	human,
}

//dangerous in small groups - skilled combatants
var mediumCharacterTypes = []string{
	animal,
	human,
	human,
	human,
	human,
	leader,
}

//as dangerous as PCs. Leader of Low or Fodder threats, unique abilities
var highCharacterTypes = []string{
	human,
	leader,
	leader,
	human,
	leader,
	human,
}

//more skilled than PCs, a few unique abilities, several Fodder threats
var heroicCharacterTypes = []string{
	leader,
	leader,
	leader,
	leader,
	leader,
	human,
}

//avenger threat. highly skilled overpowering to PCs. Can destroy cities
var soloCharacterTypes = []string{
	warMachine,
	leader,
	warMachine,
	leader,
	warMachine,
	leader,
}

func getCharacterType(t string) string {
	for _, v := range characterTypes {
		if strings.ToLower(v) == strings.ToLower(t) {
			return v
		}
	}
	return unknownRace
}

func rollCharacterType(c *character) {
	if len(c.characterType) < 1 {
		switch c.threat.name {
		case fodderThreat.name:
			c.characterType = fodderCharacterTypes[utils.Pick(fodderCharacterTypes)]
			break
		case lowThreat.name:
			c.characterType = lowCharacterTypes[utils.Pick(lowCharacterTypes)]
			break
		case mediumThreat.name:
			c.characterType = mediumCharacterTypes[utils.Pick(mediumCharacterTypes)]
			break
		case highThreat.name:
			c.characterType = highCharacterTypes[utils.Pick(highCharacterTypes)]
			break
		case heroicThreat.name:
			c.characterType = heroicCharacterTypes[utils.Pick(heroicCharacterTypes)]
			break
		case soloThreat.name:
			c.characterType = soloCharacterTypes[utils.Pick(soloCharacterTypes)]
			break
		}
	}
}
