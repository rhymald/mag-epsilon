package balance 

import (
	"rhymald/mag-epsilon/balance/common"
	// "fmt"
	// "crypto/sha512"
  // "encoding/binary"
)

func Attributes_Vitality_FromBodyStream(body *common.Stream) float64 { 
	multiplier := common.EthalonStreamLength/float64(common.GrowStep) * common.EthalonStreamLength/float64(common.GrowStep)
	// return common.DotWeightFromStreamLen(7 + body.Mean()) * multiplier
	return (common.Cbrt(body.Mean())+1) * multiplier
}

func Attributes_PoolSizeAndResistances_FromEnergyStreams(streams []*common.Stream) (float64, map[string]float64) {
	resistances, poolsize := make(map[string]float64), 0.0
	for _, each := range streams {
		poolsize += (common.Cbrt(each.Mean())+1) * common.EthalonStreamLength/float64(common.GrowStep)
		if each.Elem() != common.Elements[0] { resistances[each.Elem()] += each.Mean()+1 }
	}
	return poolsize, resistances
}