package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	// "math"
)

type Consumables struct {
	HP int
	Pool []*common.Dot
}

func BrandNewLife() *Consumables {
	var buffer Consumables
	buffer.HP = 618
	return &buffer
}

func (state *Consumables) BurnDot() (string, float64) {
	if len((*state).Pool) == 0 { return common.Elements[0], 0 }
	dot := *&state.Pool[0]
	(*state).Pool = state.Pool[1:len((*state).Pool)]
	return dot.Elem(), dot.Weight()
}
func (state *Consumables) GetDotFrom(stream *common.Stream, atts *Attributes) {
	if len((*state).Pool) > common.ChancedRound(*&atts.Poolsize) && atts.IsNPC == false { common.Wait(4096) ; return }
	var dot *common.Dot 
	if atts.IsNPC == false { 
		dot = stream.EmitDot() 
		if dot.Weight() != 0 { (*state).Pool = append((*state).Pool, dot) }
		state.Heal(1)// common.ChancedRound( 1000 / *&atts.Vitality ) ) 
		common.Wait(318)
	} else {
		state.Heal(1)// common.ChancedRound( 1000 / *&atts.Vitality ) ) 
		common.Wait(318)
	}
}
func (state *Consumables) Heal(hp int) { 
	*&state.HP = *&state.HP + hp 
	if *&state.HP < 0 { *&state.HP = 0 }  
	if *&state.HP > 1000 { *&state.HP = 1000 }
}

