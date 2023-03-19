package common

import (
	"math"
)

type Dot map[string]int 
type Stream map[string][3]int
type Element struct {
	Is, Aka string
	//   ^ Name
	Eats, EatenBy, MadeOf []string 
	// ^    ^        ^ Native: native - penetrates easier, lucky effect
  // ^    ^ Makes stream unstable, but powerful: fuel - easiest penetration
} // ^ Makes strem weak, but chills: chill - harder penetration
// else: no - regular penetration

// ELEMENTS:
var Elements []string = []string{"◌ ", "🌪 ", "🔥", "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "} 
var ElemList []Element = []Element{
	Element{Is: Elements[0], Aka: "Common", EatenBy: []string{Elements[6]}},
	Element{Is: Elements[1], Aka: "Air",    Eats: []string{Elements[3]}, EatenBy: []string{Elements[2], Elements[7]}},
	Element{Is: Elements[2], Aka: "Fire",   Eats: []string{Elements[1]}, EatenBy: []string{Elements[4], Elements[6]}},
	Element{Is: Elements[3], Aka: "Earth",  Eats: []string{Elements[4]}, EatenBy: []string{Elements[1], Elements[7]}},
	Element{Is: Elements[4], Aka: "Water",  Eats: []string{Elements[2]}, EatenBy: []string{Elements[3], Elements[6]}},
	Element{Is: Elements[5], Aka: "Void",   Eats: Elements[:], MadeOf: []string{Elements[0]}},
	Element{Is: Elements[6], Aka: "Mallom", Eats: []string{Elements[2], Elements[4]}, EatenBy: []string{Elements[5], Elements[2]}, MadeOf: []string{Elements[2], Elements[4]}},
	Element{Is: Elements[7], Aka: "Noise",  Eats: []string{Elements[1], Elements[3]}, EatenBy: []string{Elements[1], Elements[8]}, MadeOf: []string{Elements[1], Elements[3]}},
	Element{Is: Elements[8], Aka: "Resonance", Eats: Elements[1:5], EatenBy: Elements, MadeOf: []string{Elements[2]}},
}
var Physical []string = []string{"◌ ", "🌱", "🪵", "🪨", "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored
var PhysList []Element = []Element{
	Element{Is: Physical[0], Aka: "Spirit"},
	Element{Is: Physical[1], Aka: "Flesh/Plant"}, 
	Element{Is: Physical[2], Aka: "Wood/Shell"}, 
	Element{Is: Physical[3], Aka: "Hard",   MadeOf: []string{Physical[1]}}, 
	Element{Is: Physical[4], Aka: "Forged", MadeOf: []string{Physical[1], Physical[2]}}, 
}

// DOTS
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( math.Log2(str.Vol(1))) } }
// ^ Dots creation and get dot's parameers:
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()]) }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }

// STREAMS
const equalator = 1.024 // for resonation
const segmentator = 1024.0 // for creating
func BRandNewStream(elem string, leng int) *Stream {
	c, a, d := 1/float64(leng)/float64(leng)+Rand()+Rand(), 1/float64(leng)/float64(leng)+Rand()+Rand(), 1/float64(leng)/float64(leng)+Rand()+Rand()
	for r:=0; r<leng; r++ { c, a, d = c+1/float64(leng)/float64(leng)+Rand()+Rand(), a+1/float64(leng)/float64(leng)+Rand()+Rand(), d+1/float64(leng)/float64(leng)+Rand()+Rand() }
	modifier := float64(leng) / Vector(c, a, d)
	return &Stream{ elem: [3]int{ CeilRound(c*segmentator*modifier), CeilRound(a*segmentator*modifier), CeilRound(d*segmentator*modifier) }}
} // ^ streamcreation and get its parameters: 
// basic parameters:
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func (str *Stream) Cre(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][0])/segmentator+add }
func (str *Stream) Alt(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][1])/segmentator+add }
func (str *Stream) Des(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][2])/segmentator+add }
// common params:
func (str *Stream) Vol(add float64) float64 { return str.Cre(add)*str.Alt(add)*str.Des(add)-add*add*add }
func (str *Stream) Len(add float64) float64 { return Vector(str.Cre(0),str.Alt(0),str.Des(0))+add }
func (str *Stream) Mean() float64 { return 3/(1/str.Cre(0)+1/str.Alt(0)+1/str.Des(0)) }
// complex improve:  TBD
// Ca, Cd, Cad/Cda 
// Ad, Ac, Adc/Acd 
// Da, Dc, Dac/Dca