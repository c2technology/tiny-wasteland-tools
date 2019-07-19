package main


type Archetype struct {
	Name string
	Description string
	manipulate Manipulator
}

func (a Archetype) SetArchetype(player *Player){
	player.Archetype = a
	a.manipulate(player)
}

var archetypes =[]Archetype{
	{"Normals", "", func(player *Player){
		player.HitPoints = 6
		player.maxTraits = player.maxTraits + 1
	}},
	{"Mutant", "", func(player *Player){
		player.HitPoints = 8
		player.maxMutations = 3
		player.Traits["Mutation"] = Trait{Name: "Mutation"}
	}},
	{"Scavenger", "", func(player *Player){
		player.HitPoints = 7
		player.Traits["Scavenger"] = Trait{Name:"Scavenger", Description: "You gain Advantage on Scavenge Tests"}
	}},
	{"Survivor", "", func(player *Player){
		player.HitPoints = 6
		player.Traits["Always Prepared"] = Trait{Name:"Always Prepared", Description: "You gain Advantage on Tests to find shelter or aid (but not supplies or scavenging)"}
	}},
	{"Wanderer", "", func (player *Player){
		player.HitPoints = 6
		player.Traits["Unattached"] = Trait{Name:"Unattached", Description: "Anytime anyone tries to persuade, intimidate, or threaten you, they have Disadvantage"}
	}},
	{"Crazy", "", func (player *Player){
		player.HitPoints = 6
		player.Traits["Insane"] = Trait{Name:"Insane", Description: "You gain Advantage on taking risks that would make others pause"}
	}},
	{"Fixer", "", func (player *Player){
		player.HitPoints = 6
		player.Traits["Mechanic"] = Trait{Name:"Mechanic", Description: "Once per day, you may test with Disadvantage to add one Usage Rating to a scavenged item"}
	}},
	{"Tyrants", "", func (player *Player){
		player.HitPoints = 6
		player.Traits["Commanding"] = Trait{Name:"Commanding", Description: "When attacked and you do not Evade, roll 1d6. If successful, the attack misses."}
	}},
}
	
func SetArchetype(player *Player) {
	archetype := archetypes[rando.Intn(len(archetypes))]
	archetype.SetArchetype(player)
}


