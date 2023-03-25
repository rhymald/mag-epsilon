package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	"fmt"
	"crypto/sha512"
  "encoding/binary"
)

type BasicStats struct {
	ID struct {
		Entificator int
		Last int
	}
	Streams []*common.Stream
	Body *common.Stream
} 


// CREATE
func BRandNewStats(phys string) *BasicStats {
	var buffer BasicStats
	buffer.ID.Entificator = common.Epoch()
	buffer.ID.Last = common.Epoch()
	buffer.Body = common.BRandNewStream(phys, 16)
	buffer.Body.ScaleTo(common.Round(common.EthalonStreamLength*3))
	count := common.BornLuck(buffer.ID.Entificator)
	for x:=0; x<count; x++ { buffer.SproutAStream(common.Elements[0]) ; buffer.Streams[x].Attune() }
	upgrades := (5-count) * int(common.EthalonStreamLength)/common.GrowStep
	for x:=0; x<upgrades; x++ { buffer.GrowAStream(false) }
	return &buffer
}


// LEVELING
func (stats *BasicStats) SproutAStream(elem string) {
	*&stats.Streams = append(*&stats.Streams, common.BRandNewStream(elem, common.MinEnthropy))
}

func (stats *BasicStats) GrowAStream(override bool) {
	if len(*&stats.Streams) == 0 {return} // nothing to upg
	picker := common.Epoch() % len(*&stats.Streams)
	picked := *&stats.Streams[picker]
	successChance := picked.Plus( common.GrowStep ) > common.Rand()
	if successChance || override { 
		*&stats.Streams[picker] = picked
	}
}


// READ
func (stats *BasicStats) GetID() string { 
	in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Entificator))
  pid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Last))
  sid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  pstring, sstring := fmt.Sprintf("%X", pid), fmt.Sprintf("%X", sid)
	return fmt.Sprintf("%v-%v-%v-%v", pstring[:4], pstring[119:128], sstring[:1], sstring[121:128])
}
