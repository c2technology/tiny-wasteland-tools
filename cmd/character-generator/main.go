package main

import (
	"fmt"
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
	player := character.Generate("Player")
	//TODO: Pull name from input or a name generator
	fmt.Println(fmt.Sprintf("Player Name: %s", player.Name))
	fmt.Println(fmt.Sprintf("Hit Points: %d", player.HitPoints))
	fmt.Println(fmt.Sprintf("Archetype: %s", player.Archetype.Name))
	fmt.Println(fmt.Sprintf("           %s", player.Archetype.Description))
	fmt.Println("Traits:")
	for key, val := range player.Traits {
		fmt.Println(fmt.Sprintf("  %s: %s", key, val.Description))
	}
	if len(player.Mutations) > 0 {
		fmt.Println("Mutations:")
		for key, val := range player.Mutations {
			fmt.Println(fmt.Sprintf("  %s: %s", key, val.Description))
		}
	}
	if len(player.Psionics) > 0 {
		for discipline, capabilities := range player.Psionics {
			fmt.Println(fmt.Sprintf("%s Discipline:", discipline))
			for _, capability := range capabilities {
				fmt.Println(fmt.Sprintf("  %s: %s", capability.Name, capability.Description))
			}
		}
	}
}
