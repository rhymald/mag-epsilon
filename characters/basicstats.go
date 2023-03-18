package characters 

import (
	"rhymald/mag-epsilon/common"
	"fmt"
	"crypto/sha512"
  "encoding/binary"
	"math/rand"
	"time"
)

type BasicStats struct {
	ID struct {
		Entificator int
		Last int
	}
	Streams []*common.Stream
	Body *common.Stream
} 

func BRandNewStats() *BasicStats {
	var buffer BasicStats
	buffer.ID.Entificator = common.Epoch()
	buffer.ID.Last = common.Epoch()
	count := common.BornLuck(buffer.ID.Entificator)
	for x:=0; x<count; x++ { buffer.SproutAStream(common.Elements[0]) }
	for x:=0; x<5-count; x++ { buffer.GrowAStream() }
	return &buffer
}

func (stats *BasicStats) SproutAStream(elem string) {
	*&stats.Streams = append(*&stats.Streams, common.BRandNewStream(elem, 1))
}
func (stats *BasicStats) GrowAStream() {
	if len(*&stats.Streams) == 0 {return}
	picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(*&stats.Streams))
	picked := *&stats.Streams[picker]
	for {
		new := common.BRandNewStream(picked.Elem(), common.FloorRound(picked.Len(0)+1))
		negativeCheckPass := new.Alt(0)-picked.Alt(0) > 0.001 && new.Cre(0)-picked.Cre(0) > 0.001 && new.Des(0)-picked.Des(0) > 0.001 
		if negativeCheckPass {
			*&stats.Streams[picker] = new 
			return
		}
	}
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
