package character

type psionicApplicator func(psionic, *character)

//PsionicSkill within a discipline
type psionicSkill struct {
	name        string
	description string
}

//Psionic discipline
type psionic struct {
	discipline string
	skills     []psionicSkill
	apply      psionicApplicator
	revoke     psionicApplicator
}

var applyPsionic = func(p psionic, c *character) {
	c.psionics[p.discipline] = p
}
var revokePsionic = func(p psionic, c *character) {
	delete(c.psionics, p.discipline)
}
var psionicDiscipline = []string{"Telekinesis", "Telepathy", "Biomancy", "Cryomancy", "Pyromancy"}

var telekinesis = psionic{
	"Telekinesis",
	[]psionicSkill{
		{"Blast", "Test to deal 1 damage at Range. This Test is subject to all the rules of Attack."},
		{"Hurl", "As an Action, you may move any object weighing as much as you without Testing. To Hurl violently, you must make a successful Test. To Hurl objects heavier than you, you must Test with Disadvantage."},
		{"Shatter", "Test with Disadvantage to have all enemies you can see take 1 damage."},
		{"Shield", "Test to Evade until the start of your next turn. If you choose to Test with Disadvantage, you Evade with 2d6 on your next turn if successful."},
	},
	applyPsionic,
	revokePsionic,
}
var telepathy = psionic{
	"Telepathy",
	[]psionicSkill{
		{"Communicate", "You may communicate via distances to any being you are aware of. If the begin is within sight, no Test is required. Otherwise, you must make a successful Test. If they are at great distances, you must Test with Disadvantage."},
		{"Quell", "Test to quell the negative emotions in a target. If successful, you gain Advantage on your next roll against that Target."},
		{"Timeview", "Test to gain one detail about the history of an object or location you can touch or see."},
		{"Unmake", "Test with Disadvantage to have one enemy suffer Disadvantage on all Tests until the start of your next turn."},
	},
	applyPsionic,
	revokePsionic,
}
var biomancy = psionic{
	"Biomancy",
	[]psionicSkill{
		{"Bio-Organic Shock", "Test to deal 1 damage at Range. This Test is subject to all the rules of Attack. Test with Disadvantage to deal 2 damage instead."},
		{"Enhance", "Test to gain Advantage on your next Test. You may grant this to an Ally if you Test with Disadvantage."},
		{"Fast", "Test to gain 2 additional Actions this turn. You lose 2 HP at the end of those Actions."},
		{"Heal", "Test to restore 2 HP to one target. If you test with Disadvantage, you may restore 4 HP instead."},
	},
	applyPsionic,
	revokePsionic,
}
var cryomancy = psionic{
	"Cryomancy",
	[]psionicSkill{
		{"Chill", "Test to have a single target take 1 damage and gains Disadvantage on their next Test."},
		{"Coldsnap", "Test to have everything within Close range (5 ft) suffer 1 damage."},
		{"Freeze", "Test to cause one inanimate object that is about half your size or smaller to shatter and break."},
		{"Glacial", "Test to cause one target to lose an Action on their next Turn."},
	},
	applyPsionic,
	revokePsionic,
}
var pyromancy = psionic{
	"Pyromancy",
	[]psionicSkill{
		{"Burn", "Test to deal 1 damage at Range. This Test is subject to all the rules of an Attack."},
		{"Ignite", "Test with Disadvantage to cause any object roughly your size or smaller to burst into flames. Anyone who touches those flames suffers 2 damage for the round. They must Test with Disadvantage to extinguish those flames."},
		{"Extinguish", "Test to cause any flame- or heat-based Action to cool and cease."},
		{"Combustion", "Test with Disadvantage to have everything within arms' reach (or one zone) of you take 3 damage. You take 1 damage."},
	},
	applyPsionic,
	revokePsionic,
}
var psionicsTable = map[string]psionic{
	telekinesis.discipline: telekinesis,
	telepathy.discipline:   telepathy,
	biomancy.discipline:    biomancy,
	cryomancy.discipline:   cryomancy,
	pyromancy.discipline:   pyromancy,
}
