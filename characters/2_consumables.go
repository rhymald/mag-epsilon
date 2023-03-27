package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	"sync"
)

type Consumables struct {
	HP int
	Pool []*common.Dot
	Flocks []*common.Flock
	sync.Mutex
}

// CREATE
func BrandNewLife(streams int) *Consumables {
	var buffer Consumables
	buffer.HP = 618
	buffer.Flocks = append(buffer.Flocks, common.DefaultFlock(streams))
	return &buffer
}


// DOTS
func (state *Consumables) BurnDot() (string, float64) {
	if len((*state).Pool) == 0 { return common.Elements[0], 0 }
	dot := *&state.Pool[0]
	state.Lock()
	(*state).Pool = state.Pool[1:len((*state).Pool)]
	state.Unlock()
	return dot.Elem(), dot.Weight()
}

func (state *Consumables) GainDotFrom(stream *common.Stream) {
	dot := stream.EmitDot() 
	if dot.Weight() != 0 { 
		state.Lock()
		(*state).Pool = append((*state).Pool, dot) 
		state.Unlock()
	}
}


// HP
func (state *Consumables) Heal(hp int) { 
	state.Lock()
	*&state.HP = *&state.HP + hp 
	if *&state.HP < 0 { *&state.HP = 0 }  
	if *&state.HP > 1000 { *&state.HP = 1000 }
	state.Unlock()
}

