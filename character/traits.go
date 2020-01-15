package character

import (
	"fmt"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

var applyTrait = func(t trait, c *character) {
	c.traits[t.name] = t
}

type traitApplicator func(trait, *character)

var revokeTrait = func(t trait, c *character) {
	delete(c.traits, t.name)
}

type trait struct {
	name        string
	description string
	apply       traitApplicator
	revoke      traitApplicator
}

var traits = []trait{
	{"Acrobat", "You gain an Advantage when Testing to do acrobatic tricks", applyTrait, revokeTrait},
	{"Ambush Specialist", "You gain Advantage on Tests to locate, disarm, and detect ambushes and traps. You also gain Advantage on Save Tests to avoid traps.", applyTrait, revokeTrait},
	{"Armor Master", "You have 3 extra Hit Points when wearing Armor of any type. These cannot be healed until repaired (8 hours)", applyTrait, revokeTrait},
	{"Barfighter", "Your forego your Weapon Mastery and are instead proficient in Improvised Weapon. When fighting with an Improvised Weapon, you gain one additional action each turn.", func(t trait, c *character) {
		c.proficiency = proficiency{"Improvised Weapon", noop}
		delete(c.traits, t.name)
	}, revokeTrait},
	{"Beastspeaker", "You are able to communicate with animals in a primitive and simplistic manner.", applyTrait, revokeTrait},
	{"Berserker", "When attacking with a Melee Weapon, you can choose to make an attack with Disadvantage to deal 2 damage instead of 1 if you succeed.", applyTrait, revokeTrait},
	{"Blacksmith", "Once per day, you can make a Test with Advantage on any object to restore 1 Usage Rating.", applyTrait, revokeTrait},
	{"Brawler", "You evade with 2d6 when fighting Unarmed.", applyTrait, revokeTrait},
	{"Charismatic", "You gain Advantage when attempting to convince or influence someone.", applyTrait, revokeTrait},
	{"Cleave", "If your attack reduces an enemy's Hit Points to 0, you may immediately make an extra attack with Disadvantage.", applyTrait, revokeTrait},
	{"Dark-fighter", "You do not suffer Disadvantage for having your sight impaired.", applyTrait, revokeTrait},
	{"Defender", "When adjacent to an ally, you may choose to have an attack hit you before Evade Tests are made.", applyTrait, revokeTrait},
	{"Diehard", "Once per day, you may have damage that would reduce your Hit Points to 0 reduce you to 1 instead.", applyTrait, revokeTrait},
	{"Drunken Master", "While intoxicated, you may Evade without spending an Action. Additionally, you have a Disadvantage on all rolls that require delicate manipulation, social grade, or other actions that may be severely impacted by intoxication.", applyTrait, revokeTrait},
	{"Dungeoneer", "You gain Advantage when attempting to find your way through a dungeon system and when attempting to identify creatures native to subterranean systems.", applyTrait, revokeTrait},
	{"Educated", "You gain Advantage when checking to see if you know specific information.", applyTrait, revokeTrait},
	{"Eidetic Memory", "You succeed on a 4, 5, or 6 when Testing to recall information you have seen or heard, even in passing.", applyTrait, revokeTrait},
	{"Fleet of Foot", "Your speed increases from 25 to 30 feet and you gain Advantage on Tests when chasing or running.", applyTrait, revokeTrait},
	{"Healer", "As an Action, you can Test 2d6 to heal a creature next to you other than yourself. On success, they regain 2 HP. This can also be used to cure poison, disease, and other non-genetic, non-science, physical ailments.", applyTrait, revokeTrait},
	{"Insightful", "You gain Advantage when Testing to determine if someone is lying.", applyTrait, revokeTrait},
	{"Lucky", "You may reroll one Test per session.", applyTrait, revokeTrait},
	{"Marksman", "When using Focus, your next Ranged Weapon attack is successful on 3 or greater.", applyTrait, revokeTrait},
	{"Martial Artist", "You may select Unarmed as a Weapon Group. You must select a martial arts style as your Mastered Weapon. If you also have Brawler, you can Focus as a free Action, once per day.", applyTrait, revokeTrait},
	{"MacGuyver", "You can create one-use items with the right pieces. This item will grant Advantage for one Test. You may never have more than 1 item created this way. You also gain Advantage when identifying unknown items.", applyTrait, revokeTrait},
	{"Nimble Fingers", "You gain Advantage when Testing to pick locks, steam, or sleight-of-hand.", applyTrait, revokeTrait},
	{"Opportunist", "You may immediately attack with Disadvantage when an enemy within range misses an attack against you.", applyTrait, revokeTrait},
	{"Perceptive", "You gain Advantage when Testing to gain information about your surroundings or find things that may be hidden. You gain this even while you sleep.", applyTrait, revokeTrait},
	{"Psionic", "You have psionic abilities. When you use these abilities, you must mae a successful Test or the Action is wasted. This trait can be selected multiple times.", func(t trait, c *character) {
		//Determine specific Psionic Discipline
		discipline := psionicDiscipline[utils.Pick(psionicDiscipline)]
		t.name = fmt.Sprintf("Psionics (%s)", discipline)
		//Psionics can exist multiple times, remove the generic entry and replace with a discipline specific one
		delete(c.traits, "Psionics")
		c.traits[t.name] = t
		//Add psionic skills to character
		c.psionics[t.name] = psionicsTable[discipline]
	}, func(t trait, c *character) {
		delete(c.traits, t.name)
		delete(c.psionics, t.name)
	}},
	{"Quartermaster", "When you roll for Usage, you can choose to reroll once per day. You must keep the second result.", applyTrait, revokeTrait},
	{"Quick Shot", "You are able to reload a Ranged Weapon and fire it in a single Action.", applyTrait, revokeTrait},
	{"Resolute", "You gain Advantage on all Save Tests.", applyTrait, revokeTrait},
	{"Shield Bearer", "While wielding a shield, Test with 2d6 on Evade instead of 1d6. You start with a Shield.", func(t trait, c *character) {
		c.inventory = append(c.inventory, "Shield")
		c.traits[t.name] = t
	}, revokeTrait},
	{"Sneaky", "You gain Advantage when Testing to hide or sneak around without others noticing you.", applyTrait, revokeTrait},
	{"Strong", "You gain Advantage when Testing to do something with brute force.", applyTrait, revokeTrait},
	{"Survivalist", "You gain Advantage when Testing to forage for food, find water, seek shelter, or create shelter in the wild.", applyTrait, revokeTrait},
	{"Tough", "You gain 2 additional HP", func(t trait, c *character) {
		c.hitPoints = c.hitPoints + 2
		c.traits[t.name] = t
	}, func(t trait, c *character) {
		c.hitPoints = c.hitPoints - 2
		delete(c.traits, t.name)
	}},
	{"Tracker", "You gain Advantage when Testing to track someone ", applyTrait, revokeTrait},
	{"Trapmaster", "You gain Advantage on Saves against and Testing to create, locate, disarming, or Saving traps.", applyTrait, revokeTrait},
	{"Vigilant", "You gain Advantage on Initiative Tests", applyTrait, revokeTrait},
}

func rollTraits(character *character) {
	i := 0
	for (len(character.mutations) + len(character.traits)) < character.maxTraits {
		i++
		if len(character.mutations) < character.maxMutations {
			rollMutation(character)
		} else {
			rollTrait(character)
		}
		if i > 200 {
			showCharacter(*character, "")
			panic("exit")
		}
	}
}

func rollTrait(character *character) {
	if len(character.traits)+len(character.mutations) < character.maxTraits {
		trait := traits[utils.Pick(traits)]
		for _, item := range character.traits {
			if item.name == trait.name {
				rollTrait(character)
				return
			}
		}
		trait.apply(trait, character)
	}
}
