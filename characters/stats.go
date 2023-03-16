package characters 

import (
	"rhymald/mag-epsilon/common"
	"fmt"
	"crypto/sha512"
  "encoding/binary"
)

type BasicStats struct {
	ID struct {
		Entificator int
		NPC bool 
		Last int
	}
	Stats *common.Stream
} // ^ used to calculate:
type Attributes struct {
	Is map[string]bool
	Vitality float64
	Poolsize float64 
	Affinity struct {
		Power float64 
		Toughness float64
	}
} // ^ used to generate and consume: 
type Consumables struct {
	XYZ [3]int
	Pool []*common.Dot
	HP int 
}

type Character struct {
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
}

func (stats *BasicStats) GetID() string { 
	in_bytes := make([]byte, 8)
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Entificator))
  pid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  binary.LittleEndian.PutUint64(in_bytes, uint64(stats.ID.Last))
  sid := fmt.Sprintf("%X", sha512.Sum512(in_bytes))
  pstring, sstring := fmt.Sprintf("%X", pid), fmt.Sprintf("%X", sid)
	return fmt.Sprintf("%v-%v-%v-%v", pstring[:4], pstring[119:128], sstring[:1], sstring[121:128])
}
func BRandNewStats(mean int, npc bool) *BasicStats {
	var buffer BasicStats
	buffer.ID.NPC = npc 
	buffer.ID.Entificator = common.Epoch()
	buffer.ID.Last = buffer.ID.Entificator
	buffer.Stats = common.BRandNewStream(mean, common.Elements[0])
	return &buffer
}