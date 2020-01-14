package character

import (
	"fmt"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

var applyMutation = func(t Trait, c *Character) {
	c.Mutations[t.Name] = t
}

var revokeMutation = func(t Trait, c *Character) {
	delete(c.Mutations, t.Name)
}

var mutations = []Trait{
	{"Freaky Quick Reflexes", "You may reroll failed Evade Tests.", applyMutation, revokeMutation},
	{"Genetic Memory", "You may reroll at a Disadvantage when Testing to see if you know something.", applyMutation, revokeMutation},
	{"Environmental Camo", "All Tests to locate you when you are hidden are at a Disadvantage", applyMutation, revokeMutation},
	{"Bulging Muscles", "Your melee attacks do +1 damage. You gain Advantage whem Testing to lift, carry, or move something", applyMutation, revokeMutation},
	{"Third Eye", "You may reroll a a Disadvantage a failed Perception Test", applyMutation, revokeMutation},
	{"Jumpin' Jack", "You gain Advantage on any Test related to jumping, running, or moving around.", applyMutation, revokeMutation},
	{"Bone Spines", "You can protrude Bone Spines as a melee or ranged weapon. It costs an Action to deploy. You gain Advantage on the first attack each combat with this weapon. Counts as both Light Melee and Ranged. The Ranged version has an Ammo of 2 and automatically refills to 2 each day.", applyMutation, revokeMutation},
	{"Scales and Stuff", "You gain +2 HP. If you have the Diehard Trait, you can use it one additional time per day.", func(t Trait, c *Character) {
		c.HitPoints = c.HitPoints + 2
		c.Mutations[t.Name] = t
	}, func(t Trait, c *Character) {
		c.HitPoints = c.HitPoints - 2
		delete(c.Mutations, t.Name)
	}},
}

func rollMutation(character *Character) {
	fmt.Println("rolling mutation...")
	mutation := mutations[utils.Pick(mutations)]
	mutation.apply(mutation, character)
}
