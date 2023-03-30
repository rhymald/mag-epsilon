package common

import (
	// "math" // for dot only
	// "math/rand"
	// "time"
	"strconv"
	"fmt"
	"errors"
)

type Dot map[string]int 

const EthalonDotWeight float64 = 1 / 0.0132437 // min entropy
var MinWeight int = FloorRound( DotWeightFromStreamLen( BaseStreamLength/EthalonStreamLength +1) / 0.0132437)-1


// NEW
func DotWeightFromStreamLen(a float64) float64 { return Popow(Log7(a)) }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): CeilRound( EthalonDotWeight* Ntrp( DotWeightFromStreamLen(str.Len()+1) ))-MinWeight } }


// READ
func (dot *Dot) Weight() float64 { return float64((*dot)[dot.Elem()]+MinWeight) / EthalonDotWeight }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }
func (dot *Dot) ToStr() string { return fmt.Sprintf("%s|%d", dot.Elem(), (*dot)[dot.Elem()] ) }
func ParseDotFromStr(str string) (*Dot, error) {
	var new Dot
	params := Split(str)
	if len(Split(str)) != 2 { return &new, errors.New("Can't parse. That is not a dot!")} 
	w, err := strconv.Atoi(params[1])
	if err != nil { return &new, err }
	new = Dot{ params[0]: w }
	return &new, nil
}
