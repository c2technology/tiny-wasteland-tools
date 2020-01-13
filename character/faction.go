package character

import (
	"github.com/c2technology/tiny-wasteland-tools/utils"
)

const agentsOfChaos = "Agents of Chaos"
const desertRangers = "Desert Rangers"
const outlaws = "Outlaws"
const merchantsGuild = "Merchants Guild"
const projectHope = "Project Hope"
const ghostSyndicate = "Ghost Syndicate"

var factions = []string{
	agentsOfChaos,
	desertRangers,
	outlaws,
	merchantsGuild,
	projectHope,
	ghostSyndicate,
}

//RollFaction for a given Character
func RollFaction(c *Character) {
	if c.Type == Animal {
		return
	}
	faction := factions[utils.Pick(factions)]
	SetFaction(c, faction)
}

//SetFaction for a given Character with the given faction
func SetFaction(character *Character, faction string) {
	character.Faction = faction
}
