package character

import (
	"fmt"
	"sort"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

const none = ""

//characterManipulator that manipulates various attributes of a Character
type characterManipulator func(*character)

var noop = func(*character) {}

type character struct {
	name           string
	hitPoints      int
	threat         threat
	level          int
	archetype      archetype
	archetypeTrait trait
	traits         map[string]trait
	mutations      map[string]trait
	psionics       map[string]psionic
	inventory      []string
	clix           int
	proficiency    proficiency
	mastery        string
	allies         []character
	faction        string
	characterType  string
	maxMutations   int
	maxTraits      int
}

//RollCharacter with given args
func RollCharacter(name string, level int, rawArchetype string, rawProficiency string, weapon string) {
	proficiency := getProficiency(rawProficiency)
	archetype := getArchetype(rawArchetype)

	character := character{
		name:          name,
		threat:        threat{},
		traits:        make(map[string]trait),
		mutations:     make(map[string]trait),
		psionics:      make(map[string]psionic),
		characterType: "player",
		faction:       "player",
		proficiency:   proficiency,
		level:         level,
		mastery:       weapon,
		maxTraits:     3,
		inventory:     []string{"Ragged sleeping bag", "Lighter", "Belt pouch", "Cracked electric lantern (72 hour charge)", "Strong cord (50 ft)", "Rations (7 days)", "Poncho"},
	}
	rollCharacterType(&character)
	rollFaction(&character)
	rollThreat(&character)
	rollProficiency(&character)
	if len(archetype.name) > 0 {
		archetype.manipulate(&character)
	} else {
		rollArchetype(&character)
	}
	rollTraits(&character)
	adjustLevel(&character)

	if len(weapon) > 0 {
		character.inventory = append(character.inventory, weapon)
	}

	showCharacter(character, "")
}

//RollEnemy with given name and threat
func RollEnemy(name string, level int, rawThreat string, rawArchetype string, rawFaction string, rawCharacterType string, rawProficiency string, weapon string) {
	proficiency := getProficiency(rawProficiency)
	threat := getThreat(rawThreat)
	faction := getFaction(rawFaction)
	characterType := getCharacterType(rawCharacterType)
	archetype := getArchetype(rawArchetype)

	character := character{
		name:          name,
		threat:        threat,
		traits:        make(map[string]trait),
		mutations:     make(map[string]trait),
		psionics:      make(map[string]psionic),
		characterType: characterType,
		faction:       faction,
		proficiency:   proficiency,
		level:         level,
		mastery:       weapon,
		maxTraits:     3,
	}
	rollCharacterType(&character)
	rollFaction(&character)
	rollThreat(&character)
	rollProficiency(&character)
	if len(archetype.name) > 0 {
		archetype.manipulate(&character)
	} else {
		rollArchetype(&character)
	}
	rollTraits(&character)
	rollAllies(&character)
	adjustLevel(&character)

	if len(weapon) > 0 {
		character.inventory = append(character.inventory, weapon)
	}

	sort.SliceStable(character.allies, func(i, j int) bool {
		return character.allies[i].characterType < character.allies[j].characterType
	})
	showCharacter(character, "")
}

func adjustLevel(character *character) {
	if character.level < 1 {
		character.level = 1
	}
	var xp = getXP(character.level)
	for x := 1; x <= xp; x++ {
		if x%6 == 0 {
			character.hitPoints++
		}
		if x%8 == 0 {
			//Not sure how to do this?
			//proficiency++
		}
		if x%10 == 0 {
			if character.maxTraits < 7 {
				character.maxTraits++
			} else {
				//Remove a random trait
				character.maxTraits = 7
				var allTraitNames = []string{}
				for traitName := range character.traits {
					allTraitNames = append(allTraitNames, traitName)
				}
				for mutationName := range character.mutations {
					allTraitNames = append(allTraitNames, mutationName)
				}

				//Remove a random trait
				var i = utils.Roll(1, len(allTraitNames)) - 1
				traitName := allTraitNames[i]
				if trait, ok := character.traits[traitName]; ok {
					trait.revoke(trait, character)
				} else if mutation, ok := character.mutations[traitName]; ok {
					mutation.revoke(mutation, character)
				}
			}
			//Add a new trait
			rollTraits(character)
		}
	}
}

func generateEnemy(threat threat) character {
	character := character{
		level:     1,
		traits:    make(map[string]trait),
		mutations: make(map[string]trait),
		psionics:  make(map[string]psionic),
	}
	setThreat(&character, threat)
	rollCharacterType(&character)
	rollProficiency(&character)
	return character
}

func getXP(level int) int {
	var l = 1
	var xp = 0
	for l < level {
		xp++
		if xp%6 == 0 || xp%8 == 0 || xp%10 == 0 {
			l++
		}
	}
	return xp
}

func showCharacter(player character, padding string) {
	fmt.Println("=================")
	if len(player.name) > 0 {
		fmt.Println(fmt.Sprintf("%sName: %s", padding, player.name))
	}
	fmt.Println(fmt.Sprintf("%sLevel: %d", padding, player.level))
	if len(player.characterType) > 0 {
		fmt.Println(fmt.Sprintf("%sType: %s", padding, player.characterType))
	}
	if len(player.faction) > 0 {
		fmt.Println(fmt.Sprintf("%sFaction: %s", padding, player.faction))
	}
	if len(player.threat.name) > 0 {
		fmt.Println(fmt.Sprintf("%sThreat: %s", padding, player.threat.name))
		fmt.Println(fmt.Sprintf("%s   %s", padding, player.threat.description))
	}
	fmt.Println(fmt.Sprintf("%sHit Points: %d", padding, player.hitPoints))
	if len(player.proficiency.name) > 0 {
		fmt.Println(fmt.Sprintf("%sWeapon Proficiency: %s", padding, player.proficiency.name))
		fmt.Println(fmt.Sprintf("%s   You gain Disadvantage for all other weapon groups", padding))
	}

	if player.characterType != animal {
		fmt.Println(fmt.Sprintf("%sWeapon Mastery: %s", padding, player.mastery))
		fmt.Println(fmt.Sprintf("%s   You gain Advantage when using this type of weapon", padding))

		if len(player.archetype.name) > 0 {
			fmt.Println(fmt.Sprintf("%sArchetype: %s", padding, player.archetype.name))
			fmt.Println(fmt.Sprintf("%s   %s", padding, player.archetype.description))
			fmt.Println(fmt.Sprintf("%sArchetype Trait: ", padding))
			fmt.Println(fmt.Sprintf("%s   %s: %s", padding, player.archetypeTrait.name, player.archetypeTrait.description))
		}
		fmt.Println(fmt.Sprintf("%sTraits (%d):", padding, player.maxTraits))
		for key, val := range player.traits {
			fmt.Println(fmt.Sprintf("%s  %s: %s", padding, key, val.description))
		}

		if len(player.mutations) > 0 {
			fmt.Println(fmt.Sprintf("%sMutations:", padding))
			for key, val := range player.mutations {
				fmt.Println(fmt.Sprintf("%s  %s: %s", padding, key, val.description))
			}
		}
		if len(player.psionics) > 0 {
			for discipline, psionic := range player.psionics {
				fmt.Println(fmt.Sprintf("%s%s Discipline:", padding, discipline))
				for _, skill := range psionic.skills {
					fmt.Println(fmt.Sprintf("%s  %s: %s", padding, skill.name, skill.description))
				}
			}
		}
	}
	if len(player.inventory) > 0 {
		fmt.Println(fmt.Sprintf("%sInventory: %d", padding, len(player.inventory)))
		for i, item := range player.inventory {
			fmt.Println(fmt.Sprintf("%s  %d: %s", padding, i+1, item))
		}
	}
	if player.clix > 0 {
		fmt.Println(fmt.Sprintf("%sClix: %d", padding, player.clix))
	}
	if len(player.allies) > 0 {
		fmt.Println(fmt.Sprintf("Allies (%d):", len(player.allies)))
		sort.SliceStable(player.allies, func(i, j int) bool {
			return player.allies[i].threat.rank < player.allies[j].threat.rank
		})
		for _, ally := range player.allies {
			showCharacter(ally, "     ")
		}
	}
}
