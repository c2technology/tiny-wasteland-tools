package main
import (
"math/rand"
"time"
)

type Trait struct {
	Name string
	Description string
	manipulate Manipulator
}

var traitsTable =[]Trait{
	{"Acrobat", "You gain an Advantage when Testing to do acrobatic tricks", noop},
	{"Ambush Specialist", "You gain Advantage on Tests to locate, disarm, and detect ambushes and traps. You also gain Advantage on Save Tests to avoid traps.", noop},
	{"Armor Master", "", noop},
	{"Armor Master", "You have 3 extra Hit Points when wearing Armor of any type. These cannot be healed until repaired (8 hours)", noop},
	{"Barfighter", "Your forego your Weapon Mastery and are instead proficient in Improvised Weapon. When fighting with an Improvised Weapon, you gain one additional action each turn.", noop},
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
	{"Educated", "", noop},
	{"Eidetic Memory", "", noop},
	{"Fleet of Foot", "", noop},
	{"Healer", "", noop},
	{"Insightful", "", noop},
	{"Lucky", "", noop},
	{"Marksman", "", noop},
	{"Martial Artist", "", noop},
	{"MacGuyver", "", noop},
	{"Nimble Fingers", "", noop},
	{"Opportunist", "", noop},
	{"Perceptive", "", noop},
	{"Psionic", "", func(player *Player) {
		discipline := psionicDiscipline[rando.Intn(len(psionicDiscipline))]
		player.Psionics[discipline] = psionicsTable[discipline]
	}},
	{"Quartermaster", "", noop},
	{"Quick Shot", "", noop},
	{"Resolute", "", noop},
	{"Shield Bearer", "", func(player *Player){
		player.Inventory = append(player.Inventory, "Shield")
	}},
	{"Sneaky", "", noop},
	{"Strong", "", noop},
	{"Survivalist", "", noop},
	{"Tough", "", func(player *Player){
		player.HitPoints = player.HitPoints + 2
	}},
	{"Tracker", "", noop},
	{"Trapmaster", "", noop},
	{"Vigilant", "", noop},
}

var mutations = []Trait{
	{"Freaky Quick Reflexes","",noop},
	{"Genetic Memory","",noop},
	{"Environmental Camo","",noop},
	{"Bulging Muscles","",noop},
	{"Third Eye","",noop},
	{"Jumpin' Jack","",noop},
	{"Bone Spines","",noop},
	{"Scales and Stuff","",func (player *Player){
		player.HitPoints = player.HitPoints + 2
	}},
}

var psionicDiscipline = []string{"Telekinesis","Telepathy","Biomancy", "Cryomancy", "Pyromancy"}

var psionicsTable = map[string][]Trait{
	"Telekinesis":{
		{"Blast","",noop},
		{"Hurl","",noop},
		{"Shatter","",noop},
		{"Shield","",noop},
	},
	"Telepathy": {
		{"Communicate","",noop},
		{"Quell","",noop},
		{"Timeview","",noop},
		{"Unmake","",noop},
	},
	"Biomancy": {
		{"Bio-Organic Shock","",noop},
		{"Enhance","",noop},
		{"Fast","",noop},
		{"Heal","",noop},
	},
	"Cryomancy": {
		{"Chill","",noop},
		{"Coldsnap","",noop},
		{"Freeze","",noop},
		{"Glacial","",noop},
	},
	"Pyromancy": {
		{"Burn","",noop},
		{"Ignite","",noop},
		{"Extinguish","",noop},
		{"Combustion","",noop},
	},
}

var seed = rand.NewSource(time.Now().UnixNano())
var rando = rand.New(seed)
	

// AddTraits to a Player. Traits may instead be Mutations and are randomly decided
// upon and limited to the amount of Traits and/or Mutations allowed by the Player
// Archetype.
func AddTraits(player *Player) {
	for (len(player.Mutations) + len(player.Traits)) < player.maxTraits {
		if len(player.Mutations) < player.maxMutations {
			if rando.Intn(2) == 1 {
				addMutation(player)
				continue
			}
		}
		addTrait(player)
	}
}

func addTrait(player *Player) {
		trait := traitsTable[rando.Intn(len(traitsTable))]
		player.Traits[trait.Name] = trait
		trait.manipulate(player)
}

func addMutation(player *Player) {
		mutation := mutations[rando.Intn(len(mutations))]
		player.Mutations[mutation.Name] =  mutation
		mutation.manipulate(player)
}