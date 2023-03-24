package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	"math"
)

type Attributes struct {
	IsNPC bool
	Vitality float64
	Poolsize float64
	XYZ [3]int
	Resists map[string]float64
} 


// CREATE
func (stats *BasicStats) CalculaterAttributes(isnpc bool) *Attributes {
	var buffer Attributes
	buffer.Resists = make(map[string]float64)
	buffer.IsNPC = common.Epoch() < stats.ID.Last
	buffer.Vitality = common.DotWeightFromStreamLen(10 + (*stats).Body.Mean()) * common.EthalonStreamLength/float64(common.GrowStep)
	for _, each := range *&stats.Streams { 
		buffer.Poolsize += math.Log10(1 + each.Mean()) * common.EthalonStreamLength / float64(common.GrowStep)
		if each.Elem() != common.Elements[0] { buffer.Resists[each.Elem()] += math.Log2(each.Mean()+1) }
	}
	buffer.XYZ = LoginPoints[stats.ID.Last%len(LoginPoints)]
	return &buffer
}
