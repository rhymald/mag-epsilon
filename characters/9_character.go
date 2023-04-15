package characters

import (
	"sync"
	"rhymald/mag-epsilon/balance/common"
	"math"
)

type Character struct {
	View [3]int
	Atts *Attributes
	Base *BasicStats
	Cons *Consumables
	Focus []*Character
	sync.Mutex
}


// NEW
func (stats *BasicStats) ComposeCharacter(sp *SpawnPoint, cons *Consumables) *Character {
	var buffer Character
	if (*sp).Behavior == "Player" {
		buffer.Base = stats 
	} else {
		var base BasicStats // 1 / float64((*sp).Radius)/float64((*sp).Radius)/math.Pi
		prarity := math.Sqrt(1 / float64((*sp).Radius)/float64((*sp).Radius)/math.Pi)///float64(1+len(common.Split((*sp).Size.Elem()))) )
		erarity := math.Cbrt(1 / float64((*sp).Radius)/float64((*sp).Radius)/math.Pi)///float64(1+len(common.Split((*sp).Attune.Elem()))) )
		phys := common.Split((*sp).Size.Elem())[0]
		elem := common.Split((*sp).Attune.Elem())[0]
		for p:=1; p<len(common.Split((*sp).Size.Elem()));   p++ { if common.Rand() < prarity { phys = common.Split((*sp).Size.Elem())[p] }   else { break } }
		for e:=1; e<len(common.Split((*sp).Attune.Elem())); e++ { if common.Rand() < erarity { elem = common.Split((*sp).Attune.Elem())[e] } else { break } }
		base.Body = common.BRandNewStream(phys, 17) // Entro
		base.Body.ScaleTo(3) // Len
		base.SproutAStream(true) // replace
		base.Streams[0].Attune(elem)
		base.Streams[0].ScaleTo( common.Ntrp(float64(len((*stats).Streams))*stats.MeanStr()) )
		base.ID.Entificator = common.Epoch()
		base.ID.Last = base.ID.Entificator + (*sp).Lifecycle
		buffer.Base = &base
	}
	buffer.Atts = buffer.Base.CalculaterAttributes()
	buffer.Cons = cons
	base := (*sp).XYZ 
	x, y, z := common.Rand()-common.Rand(), common.Rand()-common.Rand(), common.Rand()-common.Rand()
	mod := float64((*sp).Radius) / common.Vector(x,y,z)
	(*buffer.Cons).XYZ = [3]int{ base[0]+common.ChancedRound(x*mod), base[1]+common.ChancedRound(y*mod), base[2]+common.ChancedRound(z*mod) }
	buffer.View = [3]int{ -common.ChancedRound(x*1000/common.Vector(x,y,z)), -common.ChancedRound(y*1000/common.Vector(x,y,z)), -common.ChancedRound(z*1000/common.Vector(x,y,z)) }
	return &buffer
}


// MODIFY
func (c *Character) Beware(foe *Character) {
	if c.IsFocusedOn(foe) == false { (*c).Focus = append((*c).Focus, foe) }
}

func (c *Character) CalmDown(foe *Character) {
	(*c).Focus = []*Character{}
}


// READ
func (c *Character) Where() [3]int { c.Lock() ; coords := (*c.Cons).XYZ ; c.Unlock() ; return coords }
func (c *Character) IsFocusedOn(foe *Character) bool {
	for _, each := range (*c).Focus { if each == foe { return true } } 
	return false
}