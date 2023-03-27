package characters 

import (
	// "rhymald/mag-epsilon/balance/common"
	"rhymald/mag-epsilon/balance"
	"sync"
)

type Attributes struct {
	Vitality float64
	Poolsize float64
	XYZ [3]int
	Resists map[string]float64
	sync.Mutex // for element states
} 


// CREATE
func (stats *BasicStats) CalculaterAttributes() *Attributes {
	var buffer Attributes
	buffer.Resists = make(map[string]float64)
	buffer.Vitality = balance.Attributes_Vitality_FromBodyStream((*stats).Body)
	buffer.Poolsize, buffer.Resists = balance.Attributes_PoolSizeAndResistances_FromEnergyStreams((*stats).Streams)
	// for _, each := range *&stats.Streams { 
	// 	buffer.Poolsize += (common.Cbrt(each.Mean()) +1) * common.EthalonStreamLength/float64(common.GrowStep)
	// 	if each.Elem() != common.Elements[0] { buffer.Resists[each.Elem()] += each.Mean()+1 }
	// }
	buffer.XYZ = LoginPoints[stats.ID.Last%len(LoginPoints)] // To be moved as teleport func
	return &buffer
}
