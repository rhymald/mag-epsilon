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

func (stats *BasicStats) CalculaterAttributes(isnpc bool) *Attributes {
	var buffer Attributes
	buffer.Resists = make(map[string]float64)
	buffer.IsNPC = common.Epoch() < stats.ID.Last
	buffer.Vitality = common.EthalonStreamLength / float64(common.GrowStep) * (*&stats.Body).Mean() * (1+common.DotWeightFromStreamLen((*&stats.Body).Len(0)))
	for _, each := range *&stats.Streams { 
		buffer.Poolsize += math.Sqrt((1 + each.Mean()) * (1 + each.Mean()) * (1 + each.Mean()) / (1 + common.DotWeightFromStreamLen(each.Len(0))))
		if each.Elem() != common.Elements[0] { buffer.Resists[each.Elem()] += each.Mean() }
	}
	buffer.XYZ = LoginPoints[stats.ID.Last%len(LoginPoints)]
	return &buffer
}
