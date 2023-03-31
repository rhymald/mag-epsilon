package tx 

import (
	"fmt"
	"errors"
	"rhymald/mag-epsilon/balance/common"
	"rhymald/mag-epsilon/characters"
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
	ID string // elem|name 
	Tags []string // tags
	Start int // when start
	Char string 
	Source *characters.BasicStats 
	Result string // -name, interruptedBy, failed/successRate, fatal overheat 
	End int // when it collided
	ByWith map[int][]string // streams/flocks, fractals/schemes, tools
}


// CREATE
func NewAction(what, tags string, id *characters.BasicStats) *Action {
	var buffer Action
	if len(id.GetID()) != 24 { return &buffer }
	buffer.Source = id
	buffer.Char = id.GetID()
	buffer.ID = what
	buffer.Tags = common.Split(tags)
	buffer.ByWith = make(map[int][]string)
	buffer.Result = fmt.Sprintf("Created")
	buffer.Start = common.Epoch()
	return &buffer
}


// MODIFY
func (action *Action) Feed(key int, dot *common.Dot) {
	*&action.End = common.Epoch() - *&action.Start 
	food := *&action.ByWith 
	elfood := food[key]
	elfood = append(elfood, dot.ToStr())
	food[key] = elfood
	*&action.ByWith = food
	*&action.Result = fmt.Sprintf("%s|Feeding #%d with %s%d=%0.3f", common.Split(*&action.Result)[0], key, dot.Elem(), (*dot)[dot.Elem()], dot.Weight())
}

func (action *Action) Finish(target float64, lastStr, needed int, place, direction [3]int) {
	dots := 0
	for _, each := range *&action.ByWith { dots += len(each) }
	*&action.End = common.Epoch()
	*&action.Start = - *&action.End + *&action.Start
	drate := common.ChancedRound(common.Rand()*float64(needed))
	dtarg := common.ChancedRound( float64(dots) )
	qrate := common.ChancedRound(common.Rand()*1000)
	qtarg := common.ChancedRound(target*1000)
	// success := qrate <= qtarg && drate <= dtarg
	success := drate*qrate <= dtarg*qtarg
	str := lastStr
	if success {
		leng := common.Vector(float64(direction[0]), float64(direction[1]), float64(direction[2]))
		x,y,z := 0.0, 0.0, 0.0
		if leng != 0 { 
			x = float64(direction[0])/leng * 1000
			y = float64(direction[1])/leng * 1000
			z = float64(direction[2])/leng * 1000
		}
		label := "✅ " ; if qrate > qtarg { label = "🍀 " } ; if drate > dtarg { label = "🔋 " }
		*&action.Result = fmt.Sprintf("%s| %03d<%03d %02d<%02d/%02d | to:|%3.0f|%3.0f|%3.0f| from:|%d|%d|%d", label, qrate, qtarg, drate, dots, needed, x, y, z, place[0], place[1], place[2])
	} else {
		if drate > dtarg { *&action.Result = fmt.Sprintf("🪫 | %03d<%03d %02d<%02d/%02d | by:|%d| at:|%d|%d|%d", qrate, qtarg, drate,  dots, needed, str, place[0], place[1], place[2]) }
		if qrate > qtarg { *&action.Result = fmt.Sprintf("🎲 | %03d<%03d %02d<%02d/%02d | for:|%d| at:|%d|%d|%d", qrate, qtarg, drate, dots, needed, str, place[0], place[1], place[2]) }
	}
}

//+OVERHEAT with random directions
func (action *Action) Interrupt(bywhat string, where [3]int) {
	*&action.End = common.Epoch()
	*&action.Start = - *&action.End + *&action.Start
	if len(common.Split(bywhat)) < 1 { bywhat = "UNKNOWN" }
	*&action.Result = fmt.Sprintf("⛔️ | by:|%s| at:|%d|%d|%d", bywhat, where[0], where[1], where[2])
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
func (action *Action) Elem() string { tags := common.Split((*action).ID) ; for _, each := range tags { if common.ElemInList(each, common.Elements) {return each} } ; return "ERR:NotFound" }
func (action *Action) Succeeded() bool { return action.Valid() && ( action.Failed() == false ) }
func (action *Action) Failed() bool { result := common.Split( *&action.Result ) ;  return result[0] == "⛔️ " || result[0] == "🪫 " || result[0] == "🎲 " }
func (action *Action) HasTag(tag string) bool {
	isTag, inList := false, false
	for _, each := range TagList { if each == tag { isTag = true } }
	for _, each := range (*action).Tags { if each == tag {inList = true} }
	return isTag && inList
}

func (action *Action) Valid() bool {
	finishCheck := *&action.End > 0
	if finishCheck == false { return false }
 	result := common.Split( *&action.Result ) 
	lenCheck := len(result) == 8 || len(result) == 7 || len(result) == 10
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