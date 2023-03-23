package common 

import (
	// "strconv"
	"fmt"
)

type Action struct { 
	// TX to send
	Description string // who and what
	Start int // when start
	//
	ByWith []string // streams/flocks, fractals/schemes, tools
	//
	Result string // interruptedBy, failed/successRate, fatal overheat 
	End int // when it collided
}

func NewAction(description string) *Action {
	var buffer Action
	buffer.Description = description
	*&buffer.Start = Epoch()
	return &buffer
}

func (action *Action) Feed(str int, dot *Dot) {
	food := fmt.Sprintf("%d|%s|%.0f", str, dot.Elem(), 1000*dot.Weight()) // index is always zero, so let it be weight
	*&action.ByWith = append(*&action.ByWith, food)
}
func (action *Action) Finish(target float64, dots, needed int, place, direction [3]int) {
	*&action.End = Epoch() - *&action.Start
	rate := ChancedRound(Rand()*1000)
	targ := ChancedRound( float64(dots)/float64(needed) *1000)
	success := rate <= targ
	str := "-1"
	if len(*&action.ByWith) != 0 {
		lastFeed := (*action).ByWith[len(*&action.ByWith)-1]
		str = Split(lastFeed)[0]
	}
	if success {
		rate = ChancedRound(Rand()*1000)
		targ = ChancedRound(target*1000)
		success = rate <= targ	
		if success {
			leng := Vector(float64(direction[0]), float64(direction[1]), float64(direction[2]))
			x,y,z := 0.0, 0.0, 0.0
			if leng != 0 { 
				x = float64(direction[0])/leng * 1000
				y = float64(direction[1])/leng * 1000
				z = float64(direction[2])/leng * 1000
			}
			*&action.Result = fmt.Sprintf("%d|%d|to|%.0f|%.0f|%.0f|from|%d|%d|%d", rate, targ, x, y, z, place[0], place[1], place[2])
		} else {
			*&action.Result = fmt.Sprintf("FAILED|%d|%d|for|%s|at|%d|%d|%d", rate, targ, str, place[0], place[1], place[2])
		}
	} else {
		*&action.Result = fmt.Sprintf("RUINED|%d|%d|by|%s|at|%d|%d|%d", dots, needed, str, place[0], place[1], place[2])
	}
}
func (action *Action) Interrupt(bywhat string, where [3]int) {
	*&action.End = Epoch() - *&action.Start
	if len(Split(bywhat)) < 1 { bywhat = "UNKNOWN" }
	*&action.Result = fmt.Sprintf("INTERRUPTED|%s|%d|%d|%d", bywhat, where[0], where[1], where[2])
}

func (action *Action) Succeeded() bool { return action.Valid() && ( action.Failed() == false ) }
func (action *Action) Failed() bool { result := Split( *&action.Result ) ;  return result[0] == "INTERRUPTED" || result[0] == "RUINED" || result[0] == "FAILED" }
func (action *Action) Valid() bool {
	finishCheck := *&action.End > 0
	if finishCheck == false { return false }
 	result := Split( *&action.Result ) 
	lenCheck := len(result) == 5 || len(result) == 9 || len(result) == 10
	if lenCheck == false { return false }
	// coordCheck := false
	// interrupterCheck := false : len5, interrupt, !unknown, coord
	// ruinFailCheck := false : len6, fail/ruin, at, rate<targ, coord 
	// vectorCheck := false : len7, from, vec=1, coord
	return true
}