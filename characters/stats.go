package characters 

import (
	"rhymald/mag-epsilon/common"
	"fmt"
	"crypto/sha512"
  "encoding/binary"
	"math" // bfore balance
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
} // ^ used to calculate:
type Attributes struct {
	IsNPC bool
	Vitality float64
	Poolsize float64
	XYZ [3]int
} // ^ used to generate and consume: 
type Consumables struct {
	HP int 
	Pool []*common.Dot
}

type Character struct {
	Is struct {
		Busy bool 
		Alive bool 
		NPC bool
	}
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
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
func BRandNewStats(mean int) *BasicStats {
	var buffer BasicStats
	buffer.ID.Entificator = common.Epoch()
	buffer.ID.Last = common.Epoch()
	count := common.BornLuck(buffer.ID.Entificator)
	for x:=0; x<count; x++ { buffer.SproutAStream(common.Elements[0]) }
	for x:=0; x<5-count; x++ { buffer.GrowAStream() }
	return &buffer
}

func (stats *BasicStats) CalculaterAttributes(isnpc bool) *Attributes {
	var buffer Attributes
	buffer.IsNPC = common.Epoch() < stats.ID.Last
	buffer.Vitality = 128
	buffer.Poolsize = 8
	buffer.XYZ = [3]int{0,0,0} // closest respawn
	return &buffer
}

func (state *Consumables) BurnDot() (string, float64) {
	if len((*state).Pool) == 0 { return common.Elements[0], 0 }
	dot := *&state.Pool[0]
	(*state).Pool = state.Pool[1:len((*state).Pool)]
		common.Wait(1000 / math.Sqrt(dot.Weight()) )
		return dot.Elem(), dot.Weight()
}
func (state *Consumables) GetDotFrom(stream *common.Stream, atts *Attributes) {
	if len((*state).Pool) >= common.ChancedRound(*&atts.Poolsize) && atts.IsNPC == false { common.Wait(4096) ; return }
	var dot *common.Dot 
	if atts.IsNPC == false { 
		dot = stream.EmitDot() 
		(*state).Pool = append((*state).Pool, dot)
		state.Heal( common.ChancedRound(1000 / *&atts.Vitality * dot.Weight()) ) 
		common.Wait(1000 * math.Sqrt(dot.Weight()) )
	} else {
		state.Heal( common.ChancedRound(1000 / *&atts.Vitality) ) 
		common.Wait(1000)
	}
}
func (state *Consumables) Heal(hp int) { 
	*&state.HP = *&state.HP+hp 
	if *&state.HP > 0 { *&state.HP = 0 }  
	if *&state.HP < 1000 { *&state.HP = 1000 }
}
func BrandNewLife() *Consumables {
	var buffer Consumables
	buffer.HP = 618
	return &buffer
}