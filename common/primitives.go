package common

import (
	"math"
)

type Dot map[string]int 
type Stream map[string][3]int

var ElemList []string = []string{"Common", "Air", "Fire", "Earth", "Water", "Void", "Mallom", "Noise", "Resonance"}
var Elements []string = []string{"◌ ",     "🌪 ", "🔥",   "🪨",    "🧊",    "🌑",   "🩸",     "🎶",    "☀️ "} 
const complex = 1.024
var PhysList []string = []string{"Ghosty", "Flesh", "Wooden", "Stone", "Forged"}
var Physical []string = []string{"◌ ",     "🌱",    "🪵",     "🪨",    "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored

func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( math.Log2(2+Vector(str.Des(0),str.Alt(0),str.Cre(0))) )} }
// ^ Dots creation and get dot's parameers:
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()]) }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }

func BRandNewStream(elem string, leng int) *Stream {
	c, a, d := math.Phi+Rand()+Rand(), math.Phi+Rand()+Rand(), math.Phi+Rand()+Rand()
	for r:=0; r<leng; r++ { c, a, d = c+math.Phi+Rand()+Rand(), a+math.Phi+Rand()+Rand(), d+math.Phi+Rand()+Rand() }
	modifier := float64(leng) / Vector(c, a, d)
	return &Stream{ elem: [3]int{ CeilRound(c*1024*modifier), CeilRound(a*1024*modifier), CeilRound(d*1024*modifier) }}
} // ^ streamcreation and get its parameters: 
// basic parameters:
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func (str *Stream) Cre(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][0])/1024+add }
func (str *Stream) Alt(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][1])/1024+add }
func (str *Stream) Des(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][2])/1024+add }
// common params:
func (str *Stream) Vol(add float64) float64 { return str.Cre(add)*str.Alt(add)*str.Des(add)-add*add*add }
func (str *Stream) Len(add float64) float64 { return Vector(str.Cre(0),str.Alt(0),str.Des(0))+add }
// complex improve:  TBD
// Ca, Cd, Cad/Cda 
// Ad, Ac, Adc/Acd 
// Da, Dc, Dac/Dca