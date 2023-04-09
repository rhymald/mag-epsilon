package characters 

import (
	// "rhymald/mag-epsilon/balance/common"
	"rhymald/mag-epsilon/balance"
	"sync"
)

type Attributes struct {
	Is struct {
		Busy bool // can cast
		Regen bool // can regen, out of peace
		Active bool // logged in, or spawned
		Mobile bool // can move
	}
	Vitality float64
	Poolsize float64
	Resists map[string]float64
	sync.Mutex // for element states
} 


// CREATE
func (stats *BasicStats) CalculaterAttributes() *Attributes {
	var buffer Attributes
	buffer.Resists = make(map[string]float64)
	buffer.Vitality = balance.Attributes_Vitality_FromBodyStream((*stats).Body)
	buffer.Poolsize, buffer.Resists = balance.Attributes_PoolSizeAndResistances_FromEnergyStreams((*stats).Streams)
	return &buffer
}
