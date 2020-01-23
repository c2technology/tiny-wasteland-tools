package character

import (
	"fmt"
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

type proficiency struct {
	name       string
	manipulate characterManipulator
}

var unarmed = proficiency{
	"Unarmed",
	noop,
}

var noProficiency = proficiency{"", noop}

var lightMeleeProficiency = proficiency{"Light Melee", func(c *character) {
	weapon := lightMeleeWeapons[utils.Pick(lightMeleeWeapons)]
	c.inventory = append(c.inventory, fmt.Sprintf("%s (Light Melee)", weapon))
	c.mastery = weapon
}}

var heavyMeleeProficiency = proficiency{"Heavy Melee", func(c *character) {
	weapon := heavyMeleeWeapons[utils.Pick(heavyMeleeWeapons)]
	c.inventory = append(c.inventory, fmt.Sprintf("%s (Heavy Melee)", weapon))
	c.mastery = weapon
}}

var lightRangedProficiency = proficiency{"Light Ranged", func(c *character) {
	weapon := lightRangedWeapons[utils.Pick(lightRangedWeapons)]
	c.inventory = append(c.inventory, fmt.Sprintf("%s (Light Ranged)", weapon))
	c.mastery = weapon
}}

var heavyRangedProficiency = proficiency{"Heavy Ranged", func(c *character) {
	weapon := heavyRangedWeapons[utils.Pick(heavyRangedWeapons)]
	c.inventory = append(c.inventory, fmt.Sprintf("%s (Heavy Ranged)", weapon))
	c.mastery = weapon
}}

var proficiencies = []proficiency{
	unarmed,
	lightMeleeProficiency,
	heavyMeleeProficiency,
	lightRangedProficiency,
	heavyRangedProficiency,
}

var lightMeleeWeapons = []string{
	"Spiked Knuckles",
	"Knife",
	"Hatchet",
	"Pipe",
	"Hammer",
	"Sickle",
	"Scimitar",
	"Dagger",
}
var heavyMeleeWeapons = []string{
	"Chainsaw",
	"Axe",
	"Sword",
	"Pike",
	"Sledgehammer",
	"Barbed Wire Baseball Bat",
}
var lightRangedWeapons = []string{
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
var heavyRangedWeapons = []string{
	"Shotgun",
	"Assault Rifle",
	"Marksman Rifle",
	"Bow",
	"Crossbow",
	"Missile launcher",
}

func getProficiency(proficiency string) proficiency {
	for _, v := range proficiencies {
		if strings.ToLower(v.name) == strings.ToLower(proficiency) {
			return v
		}
	}
	return noProficiency
}

func rollProficiency(character *character) {
	if character.characterType == animal {
		return
	}
	if len(character.proficiency.name) > 0 {
		return
	}
	proficiency := proficiencies[utils.Pick(proficiencies)]
	character.proficiency = proficiency
	proficiency.manipulate(character)
}
