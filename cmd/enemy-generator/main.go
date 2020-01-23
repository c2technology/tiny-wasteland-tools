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

	level := flag.Int("level", 1, "The generated character's level")
	name := flag.String("name", "Player", "the generated character's name")
	threat := flag.String("threat", "", "The generated character's threat value")
	faction := flag.String("faction", "", "The generated character's faction affiliation")
	characterType := flag.String("type", "", "The type of generated character (human, animal, leader, war machine)")
	proficiency := flag.String("proficiency", "", "The generated character's proficiency")
	weapon := flag.String("weapon", "", "The character's starting weapon")
	archetype := flag.String("archetype", "", "The character's starting archetype")
	flag.Parse()

	character.RollEnemy(*name, *level, *threat, *archetype, *faction, *characterType, *proficiency, *weapon)
}
