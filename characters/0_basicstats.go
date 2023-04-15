package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	"fmt"
	"crypto/sha512"
  "encoding/binary"
	"sync"
)

type BasicStats struct {
	ID struct {
		Entificator int
		Last int
	}
	Streams []*common.Stream
	Body *common.Stream
	sync.Mutex
} 


// CREATE
func BRandNewStats(phys string) *BasicStats {
	var buffer BasicStats
	buffer.ID.Entificator = common.Epoch()
	buffer.ID.Last = common.Epoch()
	buffer.Body = common.BRandNewStream(phys, 16)
	buffer.Body.ScaleTo(3)
	count := common.BornLuck(buffer.ID.Entificator)
	for x:=0; x<count; x++ { buffer.SproutAStream(true) }// ; buffer.Streams[x].Attune() }
	upgrades := (5-count) * int(common.EthalonStreamLength/common.GrowStep)
	for x:=0; x<upgrades; x++ { buffer.GrowAStream(true) }
	return &buffer
}


// LEVELING: tbd from flock, not global
// Destructive. Success will depend on whole default flock 
// Samevol: Void bonus
func (stats *BasicStats) SproutAStream(override bool) {
	if true || override { 
		stats.Lock()
		*&stats.Streams = append(*&stats.Streams, common.BRandNewStream(common.Elements[0], common.MinEnthropy))
		*&stats.ID.Last = common.Epoch()
		stats.Unlock()
	}
}

// Creative. Success depends on stream direction changes
// Sameshape: Resonance bonus
func (stats *BasicStats) GrowAStream(override bool) {
	if len(*&stats.Streams) == 0 {return} // nothing to upg
	picker := common.Epoch() % len(*&stats.Streams)
	picked := *&stats.Streams[picker]
	redirectSuccess := picked.Plus( common.GrowStep/common.EthalonStreamLength ) > common.Rand()
	if ( true && redirectSuccess ) || override { 
		stats.Lock()
		*&stats.Streams[picker] = picked
		*&stats.ID.Last = common.Epoch()
		stats.Unlock()
	}
}

// Alterative. Success will depend on flock's attunement 
// Sameexp: Mallom/Noise bonus
func (stats *BasicStats) BrandAStream(override bool) {
	if len(*&stats.Streams) == 0 {return} // nothing to upg
	picker := common.Epoch() % len(*&stats.Streams)
	if true || override { 
		stats.Lock()
		picked := *&stats.Streams[picker]
		picked.Attune("")
		*&stats.Streams[picker] = picked
		*&stats.ID.Last = common.Epoch()
		stats.Unlock()
	}
}


// READ
func (stats *BasicStats) IsNPC() bool { 
	if len((*stats).Streams) == 1 || (*stats).ID.Last > (*stats).ID.Entificator { return true }
	return false
}

func (stats *BasicStats) GetID() string { 
	in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Entificator))
  pid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Last))
  sid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  pstring, sstring := fmt.Sprintf("%X", pid), fmt.Sprintf("%X", sid)
	return fmt.Sprintf("%v-%v|%v-%v", pstring[:4], pstring[119:128], sstring[:1], sstring[121:128])
}

func (stats *BasicStats) MeanStr() float64 {
	meanie := 0.0
	for _, each := range (*stats).Streams { meanie += 1 / each.Len() }
	return float64(len((*stats).Streams)) / meanie
}

