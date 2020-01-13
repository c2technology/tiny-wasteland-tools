package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//Proficiency type
type Proficiency struct {
	Name       string
	manipulate characterManipulator
}

//Unarmed Proficiency
var Unarmed = Proficiency{
	"Unarmed",
	noop,
}

var proficiencies = []Proficiency{
	{"Light Melee", func(c *Character) {
		weapon := lightMelee[utils.Pick(lightMelee)]
		c.Inventory = append(c.Inventory, weapon)
		c.Mastery = weapon
	}},
	{"Heavy Melee", func(c *Character) {
		weapon := heavyMelee[utils.Pick(heavyMelee)]
		c.Inventory = append(c.Inventory, weapon)
		c.Mastery = weapon
	}},
	{"Light Ranged", func(c *Character) {
		weapon := lightRanged[utils.Pick(lightRanged)]
		c.Inventory = append(c.Inventory, weapon)
		c.Mastery = weapon
	}},
	{"Heavy Ranged", func(c *Character) {
		weapon := heavyRanged[utils.Pick(heavyRanged)]
		c.Inventory = append(c.Inventory, weapon)
		c.Mastery = weapon
	}},
}

var lightMelee = []string{
	"Spiked Knuckles",
	"Knife",
	"Hatchet",
	"Pipe",
	"Hammer",
	"Sickle",
	"Scimitar",
}
var heavyMelee = []string{
	"Chainsaw",
	"Fireman's Axe",
	"Sword",
	"Pike",
	"Sledgehammer",
	"Barbedwire Baseball Bat",
}
var lightRanged = []string{
	"Revolver",
	"Pistol",
	"Sawed-off Shotgun",
	"Hand crossbow",
	"Sling",
	"Throwing Knives",
}
var heavyRanged = []string{
	"Shotgun",
	"Assault Rifle",
	"Marksman Rifle",
	"Compound Bow",
	"Crossbow",
}

//RollProficiency for given Character
func RollProficiency(character *Character) {
	if character.Type == Animal {
		return
	}
	if len(character.Proficiency.Name) > 0 {
		return
	}
	proficiency := proficiencies[utils.Pick(proficiencies)]
	character.Proficiency = proficiency
	character.Proficiency = proficiency
	proficiency.manipulate(character)
}
