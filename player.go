package main

type Manipulator func(*Player)
var noop = func(*Player){}

type Player struct {
	Name string
	HitPoints int
	Archetype Archetype
	Traits map[string]Trait
	Mutations map[string]Trait
	Psionics map[string][]Trait
	Inventory []string
	Clix int
	Proficiency string
	Mastery string
	maxMutations int
	maxTraits int
}
func GeneratePlayer(name string) Player {
	player := Player{
		Name:         name,
		maxTraits:    3,
		maxMutations: 1,
		Traits:       make(map[string]Trait),
		Mutations:    make(map[string]Trait),
		Psionics:     make(map[string][]Trait),
	}
	SetArchetype(&player)
	AddTraits(&player)
	//	setInventory(&player)
	//	setClix(&player)
	//	setProficiency(&player)
	//	setMastery(&player)
	return player
}