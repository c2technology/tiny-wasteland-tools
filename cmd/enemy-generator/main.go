package main

import (
	"flag"

	"github.com/c2technology/tiny-wasteland-tools/character"
)

// main start of the program. This function takes optional inputs for character
// generation and performs the following:
//	selects the character archetype
//	selects traits (and optional mutant trait)
//	selects mutant trait, if applicable
//  selects weapon proficiency
//	selects weapon mastery
//	assigns starting gear
//	selects gear from a gear table
//	assigns an amount of credits
//	assigns a drive
func main() {

	levelArg := flag.Int("level", 1, "The generated character's level")
	nameArg := flag.String("name", "Player", "the generated character's name")
	threatArg := flag.String("threat", "", "The generated character's threat value")
	factionArg := flag.String("faction", "", "The generated character's faction affiliation")
	typeArg := flag.String("type", "", "The type of generated character (human, animal, leader, war machine)")
	proficiencyArg := flag.String("proficiency", "", "The generated character's proficiency")
	weaponArg := flag.String("weapon", "", "The character's starting weapon")
	flag.Parse()

	level := *levelArg
	name := *nameArg
	proficiency := character.GetProficiency(*proficiencyArg)
	weapon := *weaponArg
	threat := character.GetThreat(*threatArg)
	faction := character.GetFaction(*factionArg)
	typ := character.GetType(*typeArg)
	player := character.RollEnemy(name, level, threat, faction, typ, proficiency, weapon)
	character.ShowCharacter(player)
}
