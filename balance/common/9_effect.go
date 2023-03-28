package common 

import (
	// "strconv"
	// "fmt"
)

type Effect struct {
	// From string // who sent = [:24]
	Where [3]int // where collision happaned
	Action int // action id, when it started
	Time int
	Effect map[string]map[string]int
}


// CREATE
func (a *Action) NewEffect() (*Effect, error) {
	var buffer Effect 
	buffer.Action = (*a).End
	buffer.Time = Epoch()
	gape := float64(buffer.Time-buffer.Action) / 1000000000
	coords, direction, err := a.Where()
	for x:=0; x<3; x++ { 
		buffer.Where[x] = ChancedRound( float64(coords[x])+direction[x]*float64(gape) ) 
	} 
	return &buffer, err
}