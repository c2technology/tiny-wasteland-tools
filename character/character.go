package character

import (
	"fmt"
	"sort"
)

//None of any options. Use this as a default value when determining user input.
var None = ""

var noop = func(*Character) {}

//Character contains defined attributes
type Character struct {
	Name         string
	HitPoints    int
	Threat       Threat
	Level        int
	Archetype    Archetype
	Traits       map[string]Trait
	Mutations    map[string]Trait
	Psionics     map[string][]Trait
	Inventory    []string
	Clix         int
	Proficiency  Proficiency
	Mastery      string
	Allies       []Character
	Faction      string
	Type         string
	maxMutations int
	maxTraits    int
}

//RollPlayer Character with random attributes and the given name
func RollPlayer(name string) Character {
	character := Character{
		Name:         name,
		Level:        1,
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

//GenerateEnemy with given threat and no Allies
func GenerateEnemy(threat Threat) Character {
	character := Character{
		Level:     1,
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
	}
	SetThreat(&character, threat)
	RollType(&character)
	RollProficiency(&character)
	return character
}

//GenerateAnimal companion with given threat
func GenerateAnimal(threat Threat) Character {
	character := Character{
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
		Type:      Animal,
	}
	SetThreat(&character, threat)
	return character
}

//RollEnemy with given name and threat
func RollEnemy(name string, level int, threat Threat, faction string, typ string) Character {
	character := Character{
		Name:      name,
		Threat:    threat,
		Traits:    make(map[string]Trait),
		Mutations: make(map[string]Trait),
		Psionics:  make(map[string][]Trait),
		Type:      typ,
		Faction:   faction,
	}

	RollThreat(&character)
	RollType(&character)
	RollFaction(&character)
	RollAllies(&character)
	RollArchetype(&character)
	RollTraits(&character)
	RollProficiency(&character)

	sort.SliceStable(character.Allies, func(i, j int) bool {
		return character.Allies[i].Type < character.Allies[j].Type
	})
	return character
}

//ShowCharacter stats
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
	if len(player.Threat.Name) > 0 {
		fmt.Println(fmt.Sprintf("%sThreat: %s", padding, player.Threat.Name))
		fmt.Println(fmt.Sprintf("%s   %s", padding, player.Threat.Description))
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
		fmt.Println(fmt.Sprintf("%s   %s", padding, player.Archetype.Description))
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
	if len(player.Allies) > 0 {
		fmt.Println(fmt.Sprintf("Allies (%d):", len(player.Allies)))
		sort.SliceStable(player.Allies, func(i, j int) bool {
			return player.Allies[i].Threat.Rank < player.Allies[j].Threat.Rank
		})
		for _, sidekick := range player.Allies {
			showCharacter(sidekick, "     ")
		}
	}
}
