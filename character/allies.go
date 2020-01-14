package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

//RollAllies for a given Character. Allies are determined by the Character Threat value.
// No allies are given to Fodder and Low Threat Characters.
// Medium Threat Characters have a 20% chance of spawning 1d4 -1 allies of Medium(40%), Low(40%), or Fodder(10%) Threat.
// High Threat Characters have an 80% chance of spawning 2d4 -1 allies of Low(20%) or Fodder(80%) Threat.
// Heroic Threat Characters have an 100% chance of spawning 4d2 allies of Medium (20%), Low(20%), or Fodder(60%) Threat.
// Solo Threat Characters have an 100% chance of spawning 1 ally of Medium (30%), Low(20%), or Fodder(50%) Threat.
func RollAllies(c *Character) {
	if len(c.Allies) > 0 {
		return
	}
	var allies int

	//Determine if an Ally should be generated
	switch c.Threat.Name {
	case Medium.Name:
		if utils.Roll(1, 100) > 80 {
			allies = utils.Roll(1, 4) - 1
			for i := 0; i < allies; i++ {
				chance := utils.Roll(1, 100)
				if chance <= 40 {
					c.Allies = append(c.Allies, GenerateEnemy(Fodder))
				} else if chance <= 80 {
					c.Allies = append(c.Allies, GenerateEnemy(Low))
				} else {
					c.Allies = append(c.Allies, GenerateEnemy(Medium))
				}
			}
		}
		break
	case High.Name:
		if utils.Roll(1, 100) > 20 {
			allies = utils.Roll(2, 4) - 1
			for i := 0; i < allies; i++ {
				chance := utils.Roll(1, 100)
				if chance <= 80 {
					c.Allies = append(c.Allies, GenerateEnemy(Fodder))
				} else {
					c.Allies = append(c.Allies, GenerateEnemy(Low))
				}
			}
		}
		break
	case Heroic.Name:
		allies = utils.Roll(4, 2)
		for i := 0; i < allies; i++ {
			chance := utils.Roll(1, 100)
			if chance <= 60 {
				c.Allies = append(c.Allies, GenerateEnemy(Fodder))
			} else if chance <= 80 {
				c.Allies = append(c.Allies, GenerateEnemy(Low))
			} else {
				c.Allies = append(c.Allies, GenerateEnemy(Medium))
			}
		}
		break
	case Solo.Name:
		chance := utils.Roll(1, 100)
		if chance <= 50 {
			c.Allies = append(c.Allies, GenerateEnemy(Fodder))
		} else if chance <= 70 {
			c.Allies = append(c.Allies, GenerateEnemy(Low))
		} else {
			c.Allies = append(c.Allies, GenerateEnemy(Medium))
		}
		break
	default:
		break
	}
}