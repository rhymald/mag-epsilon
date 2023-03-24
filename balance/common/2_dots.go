package common

import (
	"math" // for dot only
	// "math/rand"
	// "time"
	// "fmt"
)

type Dot map[string]int 

const EthalonDotWeight float64 = 1024


// NEW
func DotWeightFromStreamLen(a float64) float64 { return math.Pow(math.Log10(a), math.Log10(a)) }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): CeilRound( EthalonDotWeight* Ntrp( DotWeightFromStreamLen(str.Len(1)) )) } }


// READ
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()])/EthalonDotWeight }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }