package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

const AGENTS_OF_CHAOS = "Agents of Chaos"
const DESERT_RANGERS = "Desert Rangers"
const OUTLAWS = "Outlaws"
const MERCHANTS_GUILD = "Merchants Guild"
const PROJECT_HOPE = "Project Hope"
const GHOST_SYNDICATE = "Ghost Syndicate"

var factions = []string{
	AGENTS_OF_CHAOS,
	DESERT_RANGERS,
	OUTLAWS,
	MERCHANTS_GUILD,
	PROJECT_HOPE,
	GHOST_SYNDICATE,
}

func RollFaction(c *Character) {
	if c.Type == ANIMAL {
		return
	}
	faction := factions[utils.Pick(factions)]
	SetFaction(c, faction)
}

func SetFaction(character *Character, faction string) {
	character.Faction = faction
}
