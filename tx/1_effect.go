package tx 

import (
	// "strconv"
	"rhymald/mag-epsilon/balance/common"
	"fmt"
	"errors"
)

type Effect struct {
	Where [3]int // where collision happaned
	Action int // action id, when it started
	Time int
	Element string
	Effect map[string]int
}

var TagList []string = []string{
	"Damage", // simple direct damage
	"Blast", // blast aoe
	"Projectile", // speed
	"Penetration", // pass through phys obj chance
}


// CREATE
func (a *Action) NewEffect(reach int) (*Effect, error) {
	var buffer Effect 
	buffer.Element = a.Elem()
	buffer.Action = (*a).End
	buffer.Time = buffer.Action //mili seconds
	gape := float64(buffer.Time-buffer.Action) //mili seconds
	coords, direction, err := a.Where() //mili meters
	if err != nil { return &buffer, errors.New(fmt.Sprintf("Cant parse action cordinates: %v", err)) }
	for x:=0; x<3; x++ { buffer.Where[x] = common.FloorRound( float64(coords[x])+direction[x]*float64(gape)/1000 ) } 
	// EFFECT composition block:
	buffer.Effect = make(map[string]int)
	if err == nil { err = (&buffer).Calc_Damage(a) }
	if err == nil { err = (&buffer).Calc_Blast(a) }
	// ^ +if tags in tags
	(*a).Source = nil // action is finally useless and spent
	return &buffer, err
}


// MODIFY +projectile, penetration
func (ef *Effect) Calc_Damage(a *Action) error {
	Err := errors.New("Damage calculation failed: Action has not successfully finished")
	if a.Succeeded() == false { return Err } else { Err = nil }
	if a.HasTag("Damage") == false { return nil }
	for strIndex, everyStream := range (*a).ByWith {
		stream := (*(*a).Source).Streams[strIndex-1]
		dmg := float64( (*ef).Effect["Damage"] ) / 1000 
		for _, eachDot := range everyStream {
			dot, err := common.ParseDotFromStr(eachDot)
			if err != nil { return errors.New(fmt.Sprintf("Cant parse dot: %v", err)) } else {
				// tbd elemental collisioon, balance
				dmg += (1+stream.Des())*dot.Weight() 
			}
		}
		(*ef).Effect["Damage"] = common.CeilRound(dmg * 1000)
	}
	return nil
}

func (ef *Effect) Calc_Blast(a *Action) error {
	Err := errors.New("Blast calculation failed: Action has not successfully finished")
	if a.Succeeded() == false { return Err } else { Err = nil }
	if a.HasTag("Blast") == false { return nil }
	for strIndex, _ := range (*a).ByWith {
		stream := (*(*a).Source).Streams[strIndex-1]
		aoe := float64( (*ef).Effect["Blast"] ) / 1000 
		aoe += 1 - 1 / (1+stream.Cre()) // TBD rebalance
		(*ef).Effect["Blast"] = common.CeilRound(aoe * 1000)
	}
	return nil
}