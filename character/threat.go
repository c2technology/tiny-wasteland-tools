package character

import (
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

type threat struct {
	rank        int
	name        string
	description string
	manipulate  characterManipulator
}

var fodderThreat = threat{0, "Fodder", "Fodder enemies are people or animals that have virtually no combat ability. They can be used to throw additional enemies into combat for a more epic feel.", func(c *character) {
	c.hitPoints = c.hitPoints + 1
	c.clix = utils.Roll(1, 2).Sum - 1
	c.maxMutations = 0
	c.maxTraits = 0
}}

var lowThreat = threat{1, "Low", "Low threat enemies may represent wild animals or average criminals.", func(c *character) {
	c.hitPoints = c.hitPoints + 2
	c.clix = utils.Roll(1, 6).Sum
	c.maxMutations = 0
	c.maxTraits = 0
}}

var mediumThreat = threat{2, "Medium", "Medium threat enemies can begin to be dangerous in small groups, and can represent skilled combatants or predatory creatures.", func(c *character) {
	c.hitPoints = c.hitPoints + utils.Roll(1, 3).Sum
	c.clix = utils.Roll(2, 6).Sum
}}

var highThreat = threat{3, "High", "High threat enemies are just as dangerous as a skilled Survivor. They are often leaders of Low threat or Fodder threat enemies. Since they’re usually leaders, they often have unique abilities that bolster their minions.", func(c *character) {
	c.hitPoints = c.hitPoints + utils.Roll(1, 3).Sum
	c.maxTraits = c.maxTraits + 3
	if c.maxTraits > 7 {
		c.maxTraits = 7
	}
	c.maxMutations = c.maxMutations + 1
	c.clix = 10 + utils.Roll(1, 6).Sum
}}

var heroicThreat = threat{4, "Heroic", "Heroic threat enemies are easily more skilled than your average Survivor. Provide two or three unique abilities for Heroic enemies, and several Fodder enemies to protect them.", func(c *character) {
	c.hitPoints = c.hitPoints + 2 + utils.Roll(1, 6).Sum
	c.maxTraits = 5 + utils.Roll(1, 2).Sum
	c.maxMutations = c.maxMutations + 1
	c.clix = 25 + utils.Roll(1, 25).Sum
}}

var soloThreat = threat{5, "Solo", "Solo threats are enemies that require an entire party to engage with them. This is the realm of giant monsters, city threatening war machines, and reality-warping entities. These creatures often have a wide variety of abilities to defend themselves from attackers.", func(c *character) {
	c.hitPoints = c.hitPoints + 8 + utils.Roll(1, 6).Sum
	c.maxTraits = 7
	c.maxMutations = 1 + utils.Roll(1, 3).Sum
	c.clix = 25 + utils.Roll(1, 50).Sum
}}

var unknownThreat = threat{-1, "", "", func(*character) {}}

var threats = []threat{
	fodderThreat,
	lowThreat,
	mediumThreat,
	highThreat,
	heroicThreat,
	soloThreat,
}

func getThreat(threat string) threat {
	for _, t := range threats {
		if strings.ToLower(t.name) == strings.ToLower(threat) {
			return t
		}
	}
	return unknownThreat
}

func rollThreat(c *character) {
	if c.threat.rank < 0 {
		threat := threats[utils.Pick(threats)]
		setThreat(c, threat)
	}
}

func setThreat(character *character, threat threat) {
	character.threat = threat
	threat.manipulate(character)
}
