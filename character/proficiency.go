package character

import (
	"strings"

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

var noProficiency = Proficiency{"", noop}

//LightMelee Proficiency
var LightMelee = Proficiency{"Light Melee", func(c *Character) {
	weapon := lightMelee[utils.Pick(lightMelee)]
	c.Inventory = append(c.Inventory, weapon)
	c.Mastery = weapon
}}

//HeavyMelee Proficiency
var HeavyMelee = Proficiency{"Heavy Melee", func(c *Character) {
	weapon := heavyMelee[utils.Pick(heavyMelee)]
	c.Inventory = append(c.Inventory, weapon)
	c.Mastery = weapon
}}

//LightRanged Proficiency
var LightRanged = Proficiency{"Light Ranged", func(c *Character) {
	weapon := lightRanged[utils.Pick(lightRanged)]
	c.Inventory = append(c.Inventory, weapon)
	c.Mastery = weapon
}}

//HeavyRanged Proficiency
var HeavyRanged = Proficiency{"Heavy Ranged", func(c *Character) {
	weapon := heavyRanged[utils.Pick(heavyRanged)]
	c.Inventory = append(c.Inventory, weapon)
	c.Mastery = weapon
}}

var proficiencies = []Proficiency{
	Unarmed,
	LightMelee,
	HeavyMelee,
	LightRanged,
	HeavyRanged,
}

var lightMelee = []string{
	"Spiked Knuckles",
	"Knife",
	"Hatchet",
	"Pipe",
	"Hammer",
	"Sickle",
	"Scimitar",
	"Dagger",
}
var heavyMelee = []string{
	"Chainsaw",
	"Axe",
	"Sword",
	"Pike",
	"Sledgehammer",
	"Barbedwire Baseball Bat",
}
var lightRanged = []string{
	"Bola",
	"Blowgun",
	"Atlatl",
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
	"Bow",
	"Crossbow",
	"Missle launcher",
}

//GetProficiency defined by given string. If none can be matched, return a default empty Proficiency
func GetProficiency(proficiency string) Proficiency {
	for _, v := range proficiencies {
		if strings.ToLower(v.Name) == strings.ToLower(proficiency) {
			return v
		}
	}
	return noProficiency
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
	proficiency.manipulate(character)
}
