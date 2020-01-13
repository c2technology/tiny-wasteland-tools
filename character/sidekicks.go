package character

import "github.com/c2technology/tiny-wasteland-tools/utils"

//RollSidekicks for a given Character
func RollSidekicks(c *Character) {
	//if the character is an animal type, it should only have animal sidekicks
	switch c.Threat.Name {
	case Fodder.Name:
		rollSidekicks(c, utils.Roll(2, 6), Fodder)
		break
	case Low.Name:
		rollSidekicks(c, utils.Roll(1, 4), Low)
		break
	case Medium.Name:
		rollSidekicks(c, utils.Roll(1, 4), Medium)
		break
	case High.Name:
		rollSidekicks(c, utils.Roll(1, 3), Low)
		rollSidekicks(c, utils.Roll(1, 6), Fodder)
		break
	case Heroic.Name:
		rollSidekicks(c, utils.Roll(3, 6), Fodder)
		break
	case Solo.Name:
		rollSidekicks(c, utils.Roll(4, 6), Fodder)
		break
	}
}

func rollSidekicks(c *Character, sidekicks int, threat Threat) {
	for i := 0; i < sidekicks; i++ {
		if c.Type == Animal {
			c.Sidekicks = append(c.Sidekicks, GenerateAnimal(threat))
		} else {
			c.Sidekicks = append(c.Sidekicks, GenerateEnemy(threat))
		}
	}
}
