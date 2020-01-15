package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//rollAllies for a given Character. Allies are determined by the Character Threat value.
// No allies are given to Fodder and Low Threat Characters.
// Medium Threat Characters have a 20% chance of spawning 1d4 -1 allies of Medium(40%), Low(40%), or Fodder(10%) Threat.
// High Threat Characters have an 80% chance of spawning 2d4 -1 allies of Low(20%) or Fodder(80%) Threat.
// Heroic Threat Characters have an 100% chance of spawning 4d2 allies of Medium (20%), Low(20%), or Fodder(60%) Threat.
// Solo Threat Characters have an 100% chance of spawning 1 ally of Medium (30%), Low(20%), or Fodder(50%) Threat.
func rollAllies(c *character) {
	if len(c.allies) > 0 {
		return
	}
	var allies int

	//Determine if an Ally should be generated
	switch c.threat.name {
	case mediumThreat.name:
		if utils.Roll(1, 100) > 80 {
			allies = utils.Roll(1, 4) - 1
			for i := 0; i < allies; i++ {
				chance := utils.Roll(1, 100)
				if chance <= 40 {
					c.allies = append(c.allies, generateEnemy(fodderThreat))
				} else if chance <= 80 {
					c.allies = append(c.allies, generateEnemy(lowThreat))
				} else {
					c.allies = append(c.allies, generateEnemy(mediumThreat))
				}
			}
		}
		break
	case highThreat.name:
		if utils.Roll(1, 100) > 20 {
			allies = utils.Roll(2, 4) - 1
			for i := 0; i < allies; i++ {
				chance := utils.Roll(1, 100)
				if chance <= 80 {
					c.allies = append(c.allies, generateEnemy(fodderThreat))
				} else {
					c.allies = append(c.allies, generateEnemy(lowThreat))
				}
			}
		}
		break
	case heroicThreat.name:
		allies = utils.Roll(4, 2)
		for i := 0; i < allies; i++ {
			chance := utils.Roll(1, 100)
			if chance <= 60 {
				c.allies = append(c.allies, generateEnemy(fodderThreat))
			} else if chance <= 80 {
				c.allies = append(c.allies, generateEnemy(lowThreat))
			} else {
				c.allies = append(c.allies, generateEnemy(mediumThreat))
			}
		}
		break
	case soloThreat.name:
		chance := utils.Roll(1, 100)
		if chance <= 50 {
			c.allies = append(c.allies, generateEnemy(fodderThreat))
		} else if chance <= 70 {
			c.allies = append(c.allies, generateEnemy(lowThreat))
		} else {
			c.allies = append(c.allies, generateEnemy(mediumThreat))
		}
		break
	default:
		break
	}
}

func generateAnimal(threat threat) character {
	character := character{
		traits:        make(map[string]trait),
		mutations:     make(map[string]trait),
		psionics:      make(map[string]psionic),
		characterType: animal,
	}
	setThreat(&character, threat)
	return character
}
