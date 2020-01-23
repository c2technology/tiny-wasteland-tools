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

func getFaction(name string) string {
	for _, v := range factions {
		if strings.ToLower(v) == strings.ToLower(name) {
			return v
		}
	}
	return noFaction
}

func rollFaction(c *character) {
	if len(c.faction) < 1 {
		if c.characterType == animal {
			return
		}
		faction := factions[utils.Pick(factions)]
		setFaction(c, faction)
	}
}

func setFaction(character *character, faction string) {
	character.faction = faction
}
