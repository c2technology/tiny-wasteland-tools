package character

//characterManipulator that manipulates various attributes of a Character
type characterManipulator func(*Character)

//Archetype for a Character
type Archetype struct {
	Name        string
	Description string
	manipulate  characterManipulator
}

var characterArchetypes = []Archetype{
	{"Normals", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
	}},
	{"Mutant", "", func(character *Character) {
		character.HitPoints = 8
		character.maxMutations = 3
		character.maxTraits = character.maxTraits + 1
		character.Traits["Mutation"] = Trait{Name: "Mutation"}
	}},
	{"Scavenger", "", func(character *Character) {
		character.HitPoints = 7
		character.maxTraits = character.maxTraits + 1
		character.Traits["Scavenger"] = Trait{Name: "Scavenger", Description: "You gain Advantage on Scavenge Tests"}
	}},
	{"Survivor", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
		character.Traits["Always Prepared"] = Trait{Name: "Always Prepared", Description: "You gain Advantage on Tests to find shelter or aid (but not supplies or scavenging)"}
	}},
	{"Wanderer", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
		character.Traits["Unattached"] = Trait{Name: "Unattached", Description: "Anytime anyone tries to persuade, intimidate, or threaten you, they have Disadvantage"}
	}},
	{"Crazy", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
		character.Traits["Insane"] = Trait{Name: "Insane", Description: "You gain Advantage on taking risks that would make others pause"}
	}},
	{"Fixer", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
		character.Traits["Mechanic"] = Trait{Name: "Mechanic", Description: "Once per day, you may test with Disadvantage to add one Usage Rating to a scavenged item"}
	}},
	{"Tyrants", "", func(character *Character) {
		character.HitPoints = 6
		character.maxTraits = character.maxTraits + 1
		character.Traits["Commanding"] = Trait{Name: "Commanding", Description: "When attacked and you do not Evade, roll 1d6. If successful, the attack misses."}
	}},
}

//SetArchetype randomly selected from the default, for a Character
func SetArchetype(character *Character) {
	archetype := characterArchetypes[rando.Intn(len(characterArchetypes))]
	character.Archetype = archetype
	archetype.manipulate(character)
}
