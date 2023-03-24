package common 

import (
	// "strconv"
	// "fmt"
)

type Effect struct {
	From string // who sent = [:24]
	Where [3]int // where collision happaned
	Action int // action id, when it started
	Effect struct {
		Damage map[string]int // direct dmg and heal if negative
		AoE map[string]int // line, sector, blast, chain, penetr
	} 
}

