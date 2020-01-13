package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//Threat value for any Character
type Threat struct {
	Rank        int
	Name        string
	Description string
	manipulate  characterManipulator
}

//Fodder Threat. Typically animals or people with no combat experience.
var Fodder = Threat{0, "Fodder", "Fodder enemies are people or animals that have virtually no combat ability. They can be used to throw additional enemies into combat for a more epic feel.", func(c *Character) {
	c.HitPoints = 1
	c.Clix = utils.Roll(1, 2) - 1
}}

//Low Threat. Typically wild animals or average criminals.
var Low = Threat{1, "Low", "Low threat enemies may represent wild animals or average criminals.", func(c *Character) {
	c.HitPoints = 2
	c.Clix = utils.Roll(1, 6)
}}

//Medium Threat. Predatory animals, skilled combatants, or other characters that are dangerous in small groups
var Medium = Threat{2, "Medium", "Medium threat enemies can begin to be dangerous in small groups, and can represent skilled combatants or predatory creatures.", func(c *Character) {
	c.HitPoints = c.HitPoints + utils.Roll(1, 3)
	c.Clix = utils.Roll(2, 6)
}}

//High Threat. These are just as dangerous as skilled Surviors. Often leaders of Low Threat or Fodder Threat enemies.
var High = Threat{3, "High", "High threat enemies are just as dangerous as a skilled Survivor. They are often leaders of Low threat or Fodder threat enemies. Since they’re usually leaders, they often have unique abilities that bolster their minions.", func(c *Character) {
	c.HitPoints = c.HitPoints + utils.Roll(1, 3)
	c.maxTraits = 3
	c.maxMutations = 1
	c.Clix = 10 + utils.Roll(1, 6)
}}

//Heroic Threat. More skilled than average Survivors. Equiped with abilities.
var Heroic = Threat{4, "Heroic", "Heroic threat enemies are easily more skilled than your average Survivor. Provide two or three unique abilities for Heroic enemies, and several Fodder enemies to protect them.", func(c *Character) {
	c.HitPoints = c.HitPoints + 2 + utils.Roll(1, 6)
	c.maxTraits = 5 + utils.Roll(1, 4)
	c.maxMutations = 1
	c.Clix = 25 + utils.Roll(1, 25)
}}

//Solo Threat. Requies entire party to engage. These can threaten entire cities.
var Solo = Threat{5, "Solo", "Solo threats are enemies that require an entire party to engage with them. This is the realm of giant monsters, city-Threating war machines, and reality-warping entities. These creatures often have a wide variety of abilities to defend themselves from attackers.", func(c *Character) {
	c.HitPoints = c.HitPoints + 8 + utils.Roll(1, 6)
	c.maxTraits = 6 + utils.Roll(1, 5)
	c.maxMutations = 1 + utils.Roll(1, 3)
	c.Clix = 25 + utils.Roll(1, 50)
}}

//Threats available
var Threats = []Threat{
	Fodder,
	Low,
	Medium,
	High,
	Heroic,
	Solo,
}

//RollThreat for given Character
func RollThreat(c *Character) {
	Threat := Threats[utils.Pick(Threats)]
	SetThreat(c, Threat)
}

//SetThreat for given Character
func SetThreat(character *Character, Threat Threat) {
	character.Threat = Threat
	Threat.manipulate(character)
}