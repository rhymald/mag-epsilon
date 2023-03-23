package common

import (
	"math" // for dot only
	// "math/rand"
	// "time"
)

type Dot map[string]int 

func DotWeightFromStreamLen(a float64) float64 { return  4/3 * math.Pi * math.Pow(Log(a/2),3) }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( 1000*DotWeightFromStreamLen( str.Len(1) ) ) } }
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()])/1000 }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }