package character

import (
	"fmt"
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

type Trait struct {
	Name        string
	Description string
	manipulate  characterManipulator
}

var traits = []Trait{
	{"Acrobat", "You gain an Advantage when Testing to do acrobatic tricks", noop},
	{"Ambush Specialist", "You gain Advantage on Tests to locate, disarm, and detect ambushes and traps. You also gain Advantage on Save Tests to avoid traps.", noop},
	{"Armor Master", "You have 3 extra Hit Points when wearing Armor of any type. These cannot be healed until repaired (8 hours)", noop},
	{"Barfighter", "Your forego your Weapon Mastery and are instead proficient in Improvised Weapon. When fighting with an Improvised Weapon, you gain one additional action each turn.", func(c *Character) {
		c.Proficiency = Proficiency{"Improvised Weapon", noop}
	}},
	{"Beastspeaker", "You are able to communicate with animals in a primitive and simplistic manner.", noop},
	{"Berserker", "When attacking with a Melee Weapon, you can choose to make an attack with Disadvantage to deal 2 damage instead of 1 if you succeed.", noop},
	{"Blacksmith", "Once per day, you can make a Test with Advantage on any object to restore 1 Usage Rating.", noop},
	{"Brawler", "You evade with 2d6 when fighting Unarmed.", noop},
	{"Charismatic", "You gain Advantage whem attempting to convince or influence someone.", noop},
	{"Cleave", "If your attack reduces an enemy's Hit Points to 0, you may immediately make an extra attack with Disadvantage.", noop},
	{"Dark-fighter", "You do not suffer Disadvantage for having your sight impaired.", noop},
	{"Defender", "When adjacent to an ally, you may choose to have an attack hit you before Evade Tests are made.", noop},
	{"Diehard", "Once per day, you may have damage that would reduce your Hit Points to 0 reduce you to 1 instead.", noop},
	{"Drunken Master", "While intoxicated, you may Evade without spending an Action. Additionally, you have a Disadvantage on all rolls that reqiure delicate manipilation, social grade, or other actions that may be severely impacted by intoxication.", noop},
	{"Dungeoneer", "You gain Advantage when attempting to find your way through a dungeon system and when attempting to identify creatures native to subterranean systems.", noop},
	{"Educated", "You gain Advantage when checking to see if you know specific information.", noop},
	{"Eidetic Memory", "You succeed on a 4, 5, or 6 when Testing to recall information you have seen or heard, even in passing.", noop},
	{"Fleet of Foot", "Your speed increases from 25 to 30 feet and you gain Advantage on Tests when chasing or running.", noop},
	{"Healer", "As an Action, you can Test 2d6 to heal a creature next to you other than yourself. On success, they regain 2 HP. This can also be used to cure poison, disease, and other non-genetic, non-science, physical ailments.", noop},
	{"Insightful", "You gain Advantage when Testing to determine if someone is lying.", noop},
	{"Lucky", "You may reroll one Test per session.", noop},
	{"Marksman", "When using Focus, your next Ranged Weapon attack is successfun on 3 or greater.", noop},
	{"Martial Artist", "You may select Unarmed as a Weapon Group. You must select a martial arts style as your Mastered Weaspon. If you also have Brawler, you can Focus as a free Action, once per day.", noop},
	{"MacGuyver", "You can create one-use items with the right pieces. This item will grant Advantage for one Test. You may never have more than 1 item created this way. You also gain Advantage when identifying unknown items.", noop},
	{"Nimble Fingers", "You gain Advantage when Testing to pick locks, steam, or sleight-of-hand.", noop},
	{"Opportunist", "You may immediately attack with Disadvantage when an enemy within range misses an attack against you.", noop},
	{"Perceptive", "You gain Advantage when Testing to gain information about your surroundings or find things that may be hidden. You gain this even while you sleep.", noop},
	{"Psionic", "You have psionic abilities. When you use these abilities, you must mae a successful Test or the Action is wasted. This trait can be selected multiple times.", func(character *Character) {
		discipline := psionicDiscipline[utils.Pick(psionicDiscipline)]
		//Psionics can exist multiple times, remove the generic entry and replace with a discipline specific one
		character.Traits[fmt.Sprintf("Psionics: %s", discipline)] = character.Traits["Psionics"]
		delete(character.Traits, "Psionics")
		character.Psionics[discipline] = psionicsTable[discipline]
	}},
	{"Quartermaster", "When you roll for Usage, you can choose to reroll once per day. You must keep the second result.", noop},
	{"Quick Shot", "You are able to reload a Ranged Weapon and fire it in a single Action.", noop},
	{"Resolute", "You gain Advantage on all Save Tests.", noop},
	{"Shield Bearer", "While erilding a shield, Test with 2d6 on Evade instead of 1d6. You start with a Shield.", func(character *Character) {
		character.Inventory = append(character.Inventory, "Shield")
	}},
	{"Sneaky", "You gain Advantage when Testing to hide or sneak around without others noticing you.", noop},
	{"Strong", "You gain Advantage when Testing to do something with brute force.", noop},
	{"Survivalist", "You gain Advantage when Testing to forage for food, find water, seek shelter, or create shelter in the whild.", noop},
	{"Tough", "You gain 2 additional HP", func(character *Character) {
		character.HitPoints = character.HitPoints + 2
	}},
	{"Tracker", "You gain Adgantage when Testing to track someone ", noop},
	{"Trapmaster", "You gain Advantage on Saves against and Testing to create, locate, disarming, or Saving traps.", noop},
	{"Vigilant", "You gain Advantage on Initiative Tests", noop},
}

var mutations = []Trait{
	{"Freaky Quick Reflexes", "You may reroll failed Evade Tests.", noop},
	{"Genetic Memory", "You may reroll at a Disadvantage when Testing to see if you know something.", noop},
	{"Environmental Camo", "All Tests to locate you when you are hidden are at a Disadvantage", noop},
	{"Bulging Muscles", "Your melee attacks do +1 damage. You gain Advantage whem Testing to lift, carry, or move something", noop},
	{"Third Eye", "You may reroll a a Disadvantage a failed Perception Test", noop},
	{"Jumpin' Jack", "You gain Advantage on any Test related to jumping, running, or moving around.", noop},
	{"Bone Spines", "You can protrude Bone Spines as a melee or ranged weapon. It costs an Action to deploy. You gain Advantage on the first attack each combat with this weapon. Counts as both Light Melee and Ranged. The Ranged version has an Ammo of 2 and automatically refills to 2 each day.", noop},
	{"Scales and Stuff", "You gain +2 HP. If you have the Diehard Trait, you can use it one additional time per day.", func(character *Character) {
		character.HitPoints = character.HitPoints + 2
	}},
}

var psionicDiscipline = []string{"Telekinesis", "Telepathy", "Biomancy", "Cryomancy", "Pyromancy"}

var psionicsTable = map[string][]Trait{
	"Telekinesis": {
		{"Blast", "Test to deal 1 damage at Range. This Test is subject to all the rules of Attack.", noop},
		{"Hurl", "As an Action, you may move any object weighing as much as you without Testing. To Hurl violently, you must make a successful Test. To Hurl objects heavier than you, you must Test with Disadvantage.", noop},
		{"Shatter", "Test with Disadvantage to have all enemies you can see take 1 damage.", noop},
		{"Shield", "Test to Evade until the start of your next turn. If you choose to Test with Disadvantage, you Evade with 2d6 on your next turn if successful.", noop},
	},
	"Telepathy": {
		{"Communicate", "You may communicate via distances to any being you are aware of. If the beign is within sight, no Test is required. Otherwise, you must make a successful Test. If they are at great distances, you must Test with Disadvantage.", noop},
		{"Quell", "Test to quell the negative emotions in a target. If successful, you gain Advantage on your next roll against that Target.", noop},
		{"Timeview", "Test to gain one detail about the history of an object or location you can touch or see.", noop},
		{"Unmake", "Test with Disadvantage to have one enemy suffer Disadvantage on all Tests until the start of your next turn.", noop},
	},
	"Biomancy": {
		{"Bio-Organic Shock", "Test to deal 1 damage at Range. This Test is subject to all the rules of Attack. Test with Disadvantage to deal 2 damage instead.", noop},
		{"Enhance", "Test to gain Advantage on your next Test. You may grant this to an Ally if you Test with Disadvantage.", noop},
		{"Fast", "Test to gain 2 additional Actions this turn. You lose 2 HP at the end of those Actions.", noop},
		{"Heal", "Test to restore 2 HP to one target. If you test with Disadvantage, you may restore 4 HP instead.", noop},
	},
	"Cryomancy": {
		{"Chill", "Test to have a single target take 1 damage and gains Disadvantage on their next Test.", noop},
		{"Coldsnap", "Test to have everything within Close range (5 ft) suffer 1 damage.", noop},
		{"Freeze", "Test to cause one inanimate object that is about half your size or smaller to shatter and break.", noop},
		{"Glacial", "Test to cause one target to lose an Action on their next Turn.", noop},
	},
	"Pyromancy": {
		{"Burn", "Test to deal 1 damage at Range. This Test is subject to all the rules of an Attack.", noop},
		{"Ignite", "Test with Disadvantage to cause any object roughly your size or smaller to burst into flames. Anyone who touches those flames suffers 2 damage for the round. They must Test with Disadvantage to extinguish those flames.", noop},
		{"Extinguish", "Test to cause any flame- or heat-based Action to cool and cease.", noop},
		{"Combustion", "Test with Disadvantage to have everything within arms' reach (or one zone) of you take 3 damage. You take 1 damage.", noop},
	},
}

func RollTraits(character *Character) {
	for (len(character.Mutations) + len(character.Traits)) < character.maxTraits {
		if len(character.Mutations) < character.maxMutations {
			if utils.Roll(1, 2) == 1 {
				rollMutation(character)
				continue
			}
		}
		rollTrait(character)
	}
}

func rollTrait(character *Character) {
	trait := traits[utils.Pick(traits)]
	for _, item := range character.Traits {
		if item.Name == trait.Name {
			rollTrait(character)
			return
		}
	}
	character.Traits[trait.Name] = trait
	trait.manipulate(character)
}

func rollMutation(character *Character) {
	mutation := mutations[utils.Pick(mutations)]
	character.Mutations[mutation.Name] = mutation
	mutation.manipulate(character)
}
