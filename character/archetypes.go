package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//Archetype for a Character
type Archetype struct {
	Name        string
	Description string
	manipulate  characterManipulator
}

var characterArchetypes = []Archetype{
	{"Normals", "A run of the mill, average human. Typically, these are scavengers, farmers, or living in settlements trying to repair civilization and get on with their lives as best they can.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Normal", Description: "You get to choose an additional Trait."}
		character.maxTraits = character.maxTraits + 1
	}},
	{"Mutant", "You have been warped by whatever caused the apocalypse. You are a distorted version of mankind and tend to live in your own settlements. Mutants are not typically welcomed everywhere and ranges from hulking brutes to skinny four-armed weirdos that worship cacti", func(character *Character) {
		character.HitPoints = character.HitPoints + 8
		character.maxMutations = 3
		character.ArchetypeTrait = Trait{Name: "Mutation", Description: "You are more prone to mutations."}
	}},
	{"Scavenger", "You wander the wastes looking for lost bits of stuff and putting it to use. You are hardy and are used to the harsh life in the wastes. You are typically welcomed in settlements due to supplies you find. You don't stand out much.", func(character *Character) {
		character.HitPoints = character.HitPoints + 7
		character.ArchetypeTrait = Trait{Name: "Scavenger", Description: "You gain Advantage on Scavenge Tests"}
	}},
	{"Survivor", "You survive. That's about it. You get back up when you're knocked down. You don't quit, you just keep going. You have scars and a grim haunted look in your eyes due to all of what you have seen.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Always Prepared", Description: "You gain Advantage on Tests to find shelter or aid (but not supplies or scavenging)"}
	}},
	{"Wanderer", "You move from place to place throughout the wastes picking up the pieces as you go. You don't get attached, you just keep moving on. You look like a Scavenger, but more ragged.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Unattached", Description: "Anytime anyone tries to persuade, intimidate, or threaten you, they have Disadvantage"}
	}},
	{"Crazy", "You are a nut job that loves the apocalypse. You've been driven mad by something and you keep yourself in the center of a whirlwind of chaos, violence, and blood. You typically have tattoos, piercings, and unusual style.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Insane", Description: "You gain Advantage on taking risks that would make others pause"}
	}},
	{"Fixer", "You are one who puts pieces back together. You take stuff Scavengers bring back and use it to better the lives of those around you... or to gain power. You can always spot other fixers from the amount of grease and iron filings they wear.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Mechanic", Description: "Once per day, you may test with Disadvantage to add one Usage Rating to a scavenged item"}
	}},
	{"Tyrants", "You are most likely to take something over and force your will upon the people. Just like warlords and generals, you are stuck in a perpetual war against the world, others, or yourself. Other Tyrants are typically better dressed, tougher, and fancier than others. You seek status and look the part.", func(character *Character) {
		character.HitPoints = character.HitPoints + 6
		character.ArchetypeTrait = Trait{Name: "Commanding", Description: "When attacked and you do not Evade, roll 1d6. If successful, the attack misses."}
	}},
}

// RollArchetype for given character. This will update the given Character's Threat value and set additional attributes depending on the determined Threat
func RollArchetype(character *Character) {
	archetype := characterArchetypes[utils.Pick(characterArchetypes)]
	if character.Threat.Name == Fodder.Name || character.Threat.Name == Low.Name {
		return
	}
	character.Archetype = archetype
	archetype.manipulate(character)
}
