package character

//Character contains defined attributes
type Character struct {
	Name         string
	HitPoints    int
	Archetype    Archetype
	Traits       map[string]Trait
	Mutations    map[string]Trait
	Psionics     map[string][]Trait
	Inventory    []string
	Clix         int
	Proficiency  string
	Mastery      string
	maxMutations int
	maxTraits    int
}

//Generate a Character with random attributes and the given name
func Generate(name string) Character {
	character := Character{
		Name:         name,
		maxTraits:    3,
		maxMutations: 1,
		Traits:       make(map[string]Trait),
		Mutations:    make(map[string]Trait),
		Psionics:     make(map[string][]Trait),
	}
	SetArchetype(&character)
	SetTraits(&character)
	//	setInventory(&character)
	//	setClix(&character)
	//	setProficiency(&character)
	//	setMastery(&character)
	return character
}
