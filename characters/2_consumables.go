package characters 

import (
	"rhymald/mag-epsilon/balance/common"
	"sync"
)

type Consumables struct {
	ID struct { Last int }
	HP int
	Pool []*common.Dot
	Flocks []*common.Flock
	XYZ [3]int
	// Inbound []*common.Effect
	// Outbound []*common.Action
	sync.Mutex
}

// CREATE | LOGIN
func BrandNewLife(streams int) *Consumables {
	var buffer Consumables
	buffer.HP = 382
	buffer.Flocks = append(buffer.Flocks, common.DefaultFlock(streams))
	return &buffer
}


// DOTS
func (state *Consumables) BurnDot() common.Dot {
	if len((*state).Pool) == 0 { return common.Dot{} }
	dot := *&state.Pool[0]
	state.Lock()
	(*state).Pool = state.Pool[1:len((*state).Pool)]
	state.Unlock()
	return *dot
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

