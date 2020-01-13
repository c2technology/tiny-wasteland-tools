package character

import (
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

const noFaction = ""
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

//GetFaction by name
func GetFaction(name string) string {
	for _, v := range factions {
		if strings.ToLower(v) == strings.ToLower(name) {
			return v
		}
	}
	return noFaction
}

//RollFaction for a given Character if one is not already set
func RollFaction(c *Character) {
	if len(c.Faction) < 1 {
		if c.Type == Animal {
			return
		}
		faction := factions[utils.Pick(factions)]
		SetFaction(c, faction)
	}
}

//SetFaction for a given Character with the given faction replacing any existing one
func SetFaction(character *Character, faction string) {
	character.Faction = faction
}
