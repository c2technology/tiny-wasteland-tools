package character

import (
	"fmt"
	"sort"
)

var noop = func(*Character) {}

//Character contains defined attributes
type Character struct {
	Name         string
	HitPoints    int
	Level        Level
	Archetype    Archetype
	Traits       map[string]Trait
	Mutations    map[string]Trait
	Psionics     map[string][]Trait
	Inventory    []string
	Clix         int
	Proficiency  Proficiency
	Mastery      string
	Sidekicks    []Character
	Faction      string
	Type         string
	maxMutations int
	maxTraits    int
}

//Generate a Character with random attributes and the given name
func RollPlayer(name string) Character {
	character := Character{
		Name:         name,
		maxTraits:    3,
		maxMutations: 1,
		Clix:         10,
		Traits:       make(map[string]Trait),
		Mutations:    make(map[string]Trait),
		Psionics:     make(map[string][]Trait),
		Inventory:    []string{"Ragged sleeping bag", "Lighter", "Belt pouch", "Cracked electric lantern (72 hour charge)", "Strong cord (50 ft)", "Rations (7 days)", "Poncho"},
	}
	RollArchetype(&character)
	RollTraits(&character)
	//	setInventory(&character)
	//	setClix(&character)
	RollProficiency(&character)
	//	setMastery(&character)
	return character
}

//GenerateEnemy with given level and no sidekicks
func GenerateEnemy(level Level) Character {
	character := Character{
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
	}
	SetLevel(&character, level)
	RollType(&character)
	RollProficiency(&character)
	return character
}

//GenerateEnemy with given level and no sidekicks
func GenerateAnimal(level Level) Character {
	character := Character{
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
		Type: ANIMAL,
	}
	SetLevel(&character, level)
	return character
}

func RollEnemy() Character {
	character := Character{
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
	}
	RollLevel(&character)
	RollType(&character)
	RollFaction(&character)
	RollSidekicks(&character)
	RollArchetype(&character)
	RollTraits(&character)
	RollProficiency(&character)
	return character
}

func ShowCharacter(player Character) {
	showCharacter(player, "")
}

func showCharacter(player Character, padding string) {
	fmt.Println("=================")
	if len(player.Name) > 0 {
		fmt.Println(fmt.Sprintf("%sName: %s", padding, player.Name))
	}
	if len(player.Type) > 0 {
		fmt.Println(fmt.Sprintf("%sType: %s", padding, player.Type))
	}
	if len(player.Faction) > 0 {
		fmt.Println(fmt.Sprintf("%sFaction: %s", padding, player.Faction))
	}
	if len(player.Level.Name) > 0 {
		fmt.Println(fmt.Sprintf("%sLevel: %s", padding, player.Level.Name))
		fmt.Println(fmt.Sprintf("%s   %s", padding, player.Level.Description))
	}
	fmt.Println(fmt.Sprintf("%sHit Points: %d", padding, player.HitPoints))
	if len(player.Proficiency.Name) > 0 {
		fmt.Println(fmt.Sprintf("%sWeapon Proficiency: %s", padding, player.Proficiency.Name))
		fmt.Println(fmt.Sprintf("%s   You gain Disadvantage for all other weapon groups", padding))
	}
	if len(player.Mastery) > 0 {
		fmt.Println(fmt.Sprintf("%sWeapon Mastery: %s", padding, player.Mastery))
		fmt.Println(fmt.Sprintf("%s   You gain Advantage when using this type of weapon", padding))
	}
	if len(player.Archetype.Name) > 0 {
		fmt.Println(fmt.Sprintf("%sArchetype: %s", padding, player.Archetype.Name))
		fmt.Println(fmt.Sprintf("%s   %s", padding,  player.Archetype.Description))
		fmt.Println(fmt.Sprintf("%sTraits:", padding))
	}
	if len(player.Traits) > 0 {
		for key, val := range player.Traits {
			fmt.Println(fmt.Sprintf("%s  %s: %s", padding, key, val.Description))
		}
	}
	if len(player.Mutations) > 0 {
		fmt.Println(fmt.Sprintf("%sMutations:", padding))
		for key, val := range player.Mutations {
			fmt.Println(fmt.Sprintf("%s  %s: %s", padding, key, val.Description))
		}
	}
	if len(player.Psionics) > 0 {
		for discipline, capabilities := range player.Psionics {
			fmt.Println(fmt.Sprintf("%s%s Discipline:", padding, discipline))
			for _, capability := range capabilities {
				fmt.Println(fmt.Sprintf("%s  %s: %s", padding, capability.Name, capability.Description))
			}
		}
	}
	if len(player.Inventory) > 0 {
		fmt.Println(fmt.Sprintf("%sInventory: %d", padding, len(player.Inventory)))
		for i, item := range player.Inventory {
			fmt.Println(fmt.Sprintf("%s  %d: %s", padding, i+1, item))
		}
	}
	if player.Clix > 0 {
		fmt.Println(fmt.Sprintf("%sClix: %d", padding, player.Clix))
	}
	if len(player.Sidekicks) > 0 {
		fmt.Println(fmt.Sprintf("Sidekicks (%d):", len(player.Sidekicks)))
		sort.SliceStable(player.Sidekicks, func(i, j int) bool {
			return player.Sidekicks[i].Level.Rank < player.Sidekicks[j].Level.Rank
		})
		for _, sidekick := range player.Sidekicks {
			showCharacter(sidekick, "     ")
		}
	}
}
