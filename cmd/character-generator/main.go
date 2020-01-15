package main

import (
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

	// level := flag.Int("level", 1, "The generated character's level")
	// name := flag.String("name", "Player", "the generated character's name")
	// threat := flag.String("threat", "", "the generated character's threat")
	// flag.Parse()

	// fmt.Println("Level: ", level)
	// fmt.Println("Name:", name)

	player := character.RollPlayer("Player")
	character.ShowCharacter(player)
}
