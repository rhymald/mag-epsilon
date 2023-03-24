package characters

type Character struct {
	Is struct {
		Busy bool 
		Alive bool 
		NPC bool
	}
	Target string
	View [3]int
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
}

var LoginPoints [][3]int = [][3]int{ [3]int{4,0,0}, [3]int{9,0,0}, [3]int{1,0,0}, [3]int{7,0,0} }


// CREATE
// TBD


// MODIFY
func (char *Character) LookAt(target *Character) {
	// 
}