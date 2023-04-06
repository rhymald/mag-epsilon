package characters

import (
	"sync"
	// "rhymald/mag-epsilon/balance/common"
	// "fmt"
)

type Character struct {
	Fcsd *Focus
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
}

type Focus struct {
	Busy bool 
	Focus []*Character
	View [3]int
	XYZ [3]int
	// Alive bool
	// NPC bool
	sync.Mutex
}