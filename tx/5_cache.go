package tx

import (
	"rhymald/mag-epsilon/characters"
	"rhymald/mag-epsilon/balance/common"
)

type SubstractLife struct {
	HP int
	// Barrier map[string]int
}
type SubstractDots struct {
	BurnDots int
	RegenDots []common.Dot
}
type SubstractFlock struct {
	Exp map[int][3]int
	Heat map[string]int
}
type BaseSubstract struct {
	Body common.Stream
	Energy map[int]common.Stream
}
type SubstractXYZ struct {
	Direction [3]int
	Steps int 
	Path int
}
// type SubstractCondition struct {
	// for each: What, Length, Leftover
// }

type Progress struct {
	// min rates and timeouts
	Current *characters.Character
	PrevCons *characters.Consumables
	PrevBase *characters.BasicStats
	Substract struct {
		HP []*SubstractLife
		Dots []*SubstractDots
		Flock []*SubstractFlock
		XYZ []*SubstractXYZ
		Base []*BaseSubstract
	}
}


// NEW
func InitCharCache(pl *characters.Character) *Progress {
	var buffer Progress
	buffer.Current = pl 
	buffer.PrevBase = (*pl).Base
	buffer.PrevCons = (*pl).Cons
	return &buffer
} 
