package characters 

import (
	"rhymald/mag-epsilon/common"
)

type Attributes struct {
	IsNPC bool
	Vitality float64
	Poolsize float64
	XYZ [3]int
} // ^ used to generate and consume: 

func (stats *BasicStats) CalculaterAttributes(isnpc bool) *Attributes {
	var buffer Attributes
	buffer.IsNPC = common.Epoch() < stats.ID.Last
	buffer.Vitality = 128
	buffer.Poolsize = 8
	buffer.XYZ = [3]int{0,0,0} // closest respawn
	return &buffer
}
