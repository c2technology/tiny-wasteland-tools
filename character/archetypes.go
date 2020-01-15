package character

import (
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

type archetype struct {
	name        string
	description string
	manipulate  characterManipulator
}

var noArchetype = archetype{}

var characterArchetypes = []archetype{
	{"Normals", "A run of the mill, average human. Typically, these are scavengers, farmers, or living in settlements trying to repair civilization and get on with their lives as best they can.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Normal", description: "You get to choose an additional Trait."}
		character.maxTraits = character.maxTraits + 1
	}},
	{"Mutant", "You have been warped by whatever caused the apocalypse. You are a distorted version of mankind and tend to live in your own settlements. Mutants are not typically welcomed everywhere and ranges from hulking brutes to skinny four-armed weirdos that worship cacti", func(character *character) {
		character.hitPoints = character.hitPoints + 8
		character.maxMutations = 3
		character.archetypeTrait = trait{name: "Mutation", description: "You are more prone to mutations."}
	}},
	{"Scavenger", "You wander the wastes looking for lost bits of stuff and putting it to use. You are hardy and are used to the harsh life in the wastes. You are typically welcomed in settlements due to supplies you find. You don't stand out much.", func(character *character) {
		character.hitPoints = character.hitPoints + 7
		character.archetypeTrait = trait{name: "Scavenger", description: "You gain Advantage on Scavenge Tests"}
	}},
	{"Survivor", "You survive. That's about it. You get back up when you're knocked down. You don't quit, you just keep going. You have scars and a grim haunted look in your eyes due to all of what you have seen.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Always Prepared", description: "You gain Advantage on Tests to find shelter or aid (but not supplies or scavenging)"}
	}},
	{"Wanderer", "You move from place to place throughout the wastes picking up the pieces as you go. You don't get attached, you just keep moving on. You look like a Scavenger, but more ragged.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Unattached", description: "Anytime anyone tries to persuade, intimidate, or threaten you, they have Disadvantage"}
	}},
	{"Crazy", "You are a nut job that loves the apocalypse. You've been driven mad by something and you keep yourself in the center of a whirlwind of chaos, violence, and blood. You typically have tattoos, piercings, and unusual style.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Insane", description: "You gain Advantage on taking risks that would make others pause"}
	}},
	{"Fixer", "You are one who puts pieces back together. You take stuff Scavengers bring back and use it to better the lives of those around you... or to gain power. You can always spot other fixers from the amount of grease and iron filings they wear.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Mechanic", description: "Once per day, you may test with Disadvantage to add one Usage Rating to a scavenged item"}
	}},
	{"Tyrants", "You are most likely to take something over and force your will upon the people. Just like warlords and generals, you are stuck in a perpetual war against the world, others, or yourself. Other Tyrants are typically better dressed, tougher, and fancier than others. You seek status and look the part.", func(character *character) {
		character.hitPoints = character.hitPoints + 6
		character.archetypeTrait = trait{name: "Commanding", description: "When attacked and you do not Evade, roll 1d6. If successful, the attack misses."}
	}},
}

func getArchetype(archetype string) archetype {
	for _, a := range characterArchetypes {
		if strings.ToLower(a.name) == strings.ToLower(archetype) {
			return a
		}
	}
	return noArchetype
}

func rollArchetype(character *character) {
	archetype := characterArchetypes[utils.Pick(characterArchetypes)]
	if character.threat.name == fodderThreat.name || character.threat.name == lowThreat.name {
		return
	}
	character.archetype = archetype
	archetype.manipulate(character)
}
