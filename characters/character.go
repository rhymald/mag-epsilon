package characters

type Character struct {
	Is struct {
		Busy bool 
		Alive bool 
		NPC bool
	}
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
}
