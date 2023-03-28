package tx 

import (
	"fmt"
	"errors"
	"rhymald/mag-epsilon/balance/common"
	"strconv"
)

// Basic actions: 
//	Login, - read dots from latest logout and hp or default 
//  Logout, Die - consume dots for next login
//  Jinx, - attack 
//  Loot, - 
//  Drop, - 
//  Regen-, Move- snapshots

type Action struct { 
	Source string 
	Start int // when start
	End int // when it collided
	Result string // interruptedBy, failed/successRate, fatal overheat 
	ByWith map[string][]string // streams/flocks, fractals/schemes, tools
}


// CREATE
func NewAction(id ...string) *Action {
	var buffer Action
	if len(id[0]) != 24 { return &buffer }
	// description := fmt.Sprintf("ID=%0X", Epoch()%(256*256*256))
	// for _, each := range descriptions { description = fmt.Sprintf("%s|%s", description, each) }
	buffer.Source = id[0]
	buffer.ByWith = make(map[string][]string)
	buffer.Start = common.Epoch()
	return &buffer
}


// MODIFY
func (action *Action) Feed(key string, dot *common.Dot) {
	food := *&action.ByWith 
	elfood := food[key]
	elfood = append(elfood, fmt.Sprintf("%s|%d", dot.Elem(), (*dot)[dot.Elem()] ))
	food[key] = elfood
	*&action.ByWith = food
}

func (action *Action) Finish(target float64, dots, needed int, place, direction [3]int) {
	*&action.End = common.Epoch()
	*&action.Start = - *&action.End + *&action.Start
	rate := common.ChancedRound(common.Rand()*1000)
	targ := common.ChancedRound( float64(dots)/float64(needed) *1000)
	success := rate <= targ
	str := -1
	if success {
		rate = common.ChancedRound(common.Rand()*1000)
		targ = common.ChancedRound(target*1000)
		success = rate <= targ	
		if success {
			leng := common.Vector(float64(direction[0]), float64(direction[1]), float64(direction[2]))
			x,y,z := 0.0, 0.0, 0.0
			if leng != 0 { 
				x = float64(direction[0])/leng * 1000
				y = float64(direction[1])/leng * 1000
				z = float64(direction[2])/leng * 1000
			}
			*&action.Result = fmt.Sprintf("%d|%d|to|%.0f|%.0f|%.0f|from|%d|%d|%d", rate, targ, x, y, z, place[0], place[1], place[2])
		} else {
			*&action.Result = fmt.Sprintf("FAILED|%d|%d|for|%d|at|%d|%d|%d", rate, targ, str, place[0], place[1], place[2])
		}
	} else {
		*&action.Result = fmt.Sprintf("RUINED|%d|%d|by|%d|at|%d|%d|%d", dots, needed, str, place[0], place[1], place[2])
	}
}

func (action *Action) Interrupt(bywhat string, where [3]int) {
	*&action.End = common.Epoch()
	*&action.Start = - *&action.End + *&action.Start
	if len(common.Split(bywhat)) < 1 { bywhat = "UNKNOWN" }
	*&action.Result = fmt.Sprintf("INTERRUPTED|by|%s|at|%d|%d|%d", bywhat, where[0], where[1], where[2])
}


// READ
func (action *Action) Where() ([3]int, [3]float64, error) {
	result := common.Split( *&action.Result )
	var buffer [3]int
	direction, vec := [3]float64{0,0,0}, 0.0
	if len(result) < 3 { return buffer, direction, errors.New("Invalid result: has no coordinates") }
	xstr, ystr, zstr := result[len(result)-3], result[len(result)-2], result[len(result)-1]
	x, err := strconv.Atoi(xstr) ; if err != nil { return buffer, direction, errors.New("Invalid result: x[-3] cant be parsed") }
	y, err := strconv.Atoi(ystr) ; if err != nil { return buffer, direction, errors.New("Invalid result: y[-2] cant be parsed") }
	z, err := strconv.Atoi(zstr) ; if err != nil { return buffer, direction, errors.New("Invalid result: z[-1] cant be parsed") }
	buffer = [3]int{x,y,z}
	if true {
		astr, bstr, cstr := result[len(result)-7], result[len(result)-6], result[len(result)-5]
		a, err := strconv.Atoi(astr) ; if err != nil { direction = [3]float64{0,0,0} } else {
			b, err := strconv.Atoi(bstr) ; if err != nil { direction = [3]float64{0,0,0} } else {
				c, err := strconv.Atoi(cstr) ; if err != nil { direction = [3]float64{0,0,0} } else {
					vec = common.Vector( float64(a)/1000, float64(b)/1000, float64(c)/1000 )
					if vec != 0 { direction = [3]float64{ float64(a)/vec, float64(b)/vec, float64(c)/vec } }
				}
			}
		}
	}
	return buffer, direction, nil
}

// func (action *Action) Describe() map[string]string { return ParseTags((*action).Description) }
func (action *Action) Succeeded() bool { return action.Valid() && ( action.Failed() == false ) }
func (action *Action) Failed() bool { result := common.Split( *&action.Result ) ;  return result[0] == "INTERRUPTED" || result[0] == "RUINED" || result[0] == "FAILED" }

func (action *Action) Valid() bool {
	finishCheck := *&action.End > 0
	if finishCheck == false { return false }
 	result := common.Split( *&action.Result ) 
	lenCheck := len(result) == 7 || len(result) == 9 || len(result) == 10
	if lenCheck == false { return false }
	coordCheck, vectorCheck := false, false
	_, vector, err := action.Where() ; if err == nil { coordCheck = true } else { return false }
	if action.Failed() == false && coordCheck == true {
		// labels := result[2] == "to" && result[6] == "from"
		if ( common.Vector(vector[0],vector[1],vector[2]) >= 998 && common.Vector(vector[0],vector[1],vector[2]) <= 1002 ) || common.Vector(vector[0],vector[1],vector[2]) <= 2 { 
			vectorCheck = true } else { return false }
	} else { vectorCheck = true }
	// interrupterCheck := false : len5, interrupt, !unknown, coord
	// ruinFailCheck := false : len6, fail/ruin, at, rate<targ, coord 
	// vectorCheck := false : from
	return true && coordCheck && vectorCheck
}