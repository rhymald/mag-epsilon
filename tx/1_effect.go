package tx 

import (
	// "strconv"
	"rhymald/mag-epsilon/balance/common"
	"fmt"
	"errors"
)

type Effect struct {
	// From string // who sent = [:24]
	// +element
	Where [3]int // where collision happaned
	Action int // action id, when it started
	Time int
	Effect map[string]int
}


// CREATE
func (a *Action) NewEffect(reach int) (*Effect, error) {
	var buffer Effect 
	buffer.Action = (*a).End
	buffer.Time = common.Epoch() //mili seconds
	gape := float64(buffer.Time-buffer.Action) //mili seconds
	coords, direction, err := a.Where() //mili meters
	if err != nil { return &buffer, errors.New(fmt.Sprintf("Cant parse action cordinates: %v", err)) }
	for x:=0; x<3; x++ { buffer.Where[x] = common.FloorRound( float64(coords[x])+direction[x]*float64(gape)/1000 ) } 
	// EFFECT composition block:
	buffer.Effect = make(map[string]int)
	if err == nil { err = (&buffer).Calc_Damage(a) }
	if err == nil { err = (&buffer).Calc_Spread(a) }
	// +projectile: + time, + XYZ instead of later
	// ^ +if tags in tags
	(*a).Source = nil
	return &buffer, err
}


// MODIFY +projectile, penetration
func (ef *Effect) Calc_Damage(a *Action) error {
	Err := errors.New("Damage calculation failed: Action has not successfully finished")
	if a.Succeeded() == false { return Err } else { Err = nil }
	for strIndex, everyStream := range (*a).ByWith {
		stream := (*(*a).Source).Streams[strIndex-1]
		dmg := float64( (*ef).Effect["DMG"] ) / 1000 // TBD replce with fractal elem
		for _, eachDot := range everyStream {
			dot, err := common.ParseDotFromStr(eachDot)
			if err != nil { return errors.New(fmt.Sprintf("Cant parse dot: %v", err)) } else {
				dmg += (1+stream.Des())*dot.Weight() 
			}
		}
		(*ef).Effect["Damage"] = common.CeilRound(dmg * 1000)
	}
	return nil
}

func (ef *Effect) Calc_Spread(a *Action) error {
	Err := errors.New("Damage calculation failed: Action has not successfully finished")
	if a.Succeeded() == false { return Err } else { Err = nil }
	for strIndex, _ := range (*a).ByWith {
		stream := (*(*a).Source).Streams[strIndex-1]
		aoe := float64( (*ef).Effect["AOE"] ) / 1000 // TBD replce with fractal elem
		aoe += 1 - 1 / (1+stream.Cre()) // TBD realance
		(*ef).Effect["Spread"] = common.CeilRound(aoe * 1000)
	}
	return nil
}