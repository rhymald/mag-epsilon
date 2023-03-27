package characters

import "sync"

type Character struct {
	Fcsd *Focus
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
}

type Focus struct {
	Is struct {
		Busy bool 
		Alive bool 
		NPC bool
	}
	Target string
	View [3]int
	sync.Mutex
}

var LoginPoints [][3]int = [][3]int{ [3]int{4,0,0}, [3]int{9,0,0}, [3]int{1,0,0}, [3]int{7,0,0} }


// CREATE
// TBD


// MODIFY
func (char *Character) LookAt(target *Character) {
	// 
}