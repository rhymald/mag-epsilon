package common

import (
	"math" // for dot only
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


    ////////////// 
   // ELEMENTS // 
	//////////////
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

    //////////
   // DOTS //
  //////////
func DotWeightFromStreamLen(a float64) float64 { return  4/3 * math.Pi * math.Pow(Log(a/2),3) }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( DotWeightFromStreamLen( str.Len(1) ) ) } }
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()]) }
func (dot *Dot) Elem() string { for elem, _ := range *dot { return elem } ; return "ERR" }


    /////////////
   // STREAMS //
  /////////////
const EthalonStreamLength float64 = 1024
const BaseStreamLength float64 = 512
const GrowStep int = 128
const MinEnthropy int = 2
func BRandNewStream(elem string, length int) *Stream {
	leng := BaseStreamLength / 2
	if elem == Elements[0] { leng = BaseStreamLength }
	if elem == Elements[5] { leng = BaseStreamLength / 4 }
	if elem == Elements[8] { leng = BaseStreamLength / 8 }
	enthropy := 1/float64(length+1)/float64(length+1)
	c, a, d := (1+Rand()-Rand())*enthropy, (1+Rand()-Rand())*enthropy, (1+Rand()-Rand())*enthropy
	for step:=0; step<length-1; step++ { c, a, d = c+(1+Rand()-Rand())*enthropy, a+(1+Rand()-Rand())*enthropy, d+(1+Rand()-Rand())*enthropy }
	modifier := 1 / Vector(c, a, d)
	return &Stream{ elem: [3]int{ CeilRound(c*leng*modifier), CeilRound(a*leng*modifier), CeilRound(d*leng*modifier) }}
} 
// func (str *Stream) mean() float64 { return 3/(1/str.Cre(0)+1/str.Alt(0)+1/str.Des(0)) }
func (str *Stream) Len(add float64) float64 { return Vector(str.Cre(0),str.Alt(0),str.Des(0))+add }
func (str *Stream) Harmony() float64 {  return math.Log2(str.Len(0) / math.Max(math.Max(str.Cre(0), str.Alt(0)), str.Des(0))) / math.Log2( math.Sqrt(3) ) }
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func (str *Stream) Cre(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][0])/EthalonStreamLength+add }
func (str *Stream) Alt(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][1])/EthalonStreamLength+add }
func (str *Stream) Des(add float64) float64 { buf := *str ; return float64(buf[str.Elem()][2])/EthalonStreamLength+add }
func (str *Stream) RandShapeAs(cc, aa, dd int) {
	keepLen := str.Len(0) * EthalonStreamLength
	entroc, entroa, entrod := 1/float64(cc+1)/float64(cc+1), 1/float64(aa+1)/float64(aa+1), 1/float64(dd+1)/float64(dd+1)
	c, a, d := (1+Rand()-Rand())*entroc, (1+Rand()-Rand())*entroa, (1+Rand()-Rand())*entrod 
	for step:=0; step<cc-1; step++ { c += entroc*(1+Rand()-Rand()) }
	for step:=0; step<aa-1; step++ { a += entroa*(1+Rand()-Rand()) }
	for step:=0; step<dd-1; step++ { d += entrod*(1+Rand()-Rand()) }
	modifier := 1 / Vector(c, a, d)
	*str = Stream{ str.Elem(): [3]int{ Round(c*keepLen*modifier), Round(a*keepLen*modifier), Round(d*keepLen*modifier) }}
}
func (str *Stream) ScaleTo(ll int) {
	multiplier := float64(ll) / EthalonStreamLength / str.Len(0)
	c, a, d := str.Cre(0)*EthalonStreamLength, str.Alt(0)*EthalonStreamLength, str.Des(0)*EthalonStreamLength
	*str = Stream{ str.Elem(): [3]int{ Round(c*multiplier), Round(a*multiplier), Round(d*multiplier) }}
}
func (str *Stream) Plus(ll int) float64 {
	increasement := BRandNewStream(str.Elem(), MinEnthropy+ll/GrowStep)
	increasement.ScaleTo(ll)
	newLen := str.Len(0)*EthalonStreamLength+float64(ll)
	*str = Stream{ str.Elem(): [3]int{ 
		Round((str.Cre(0)+increasement.Cre(0))*EthalonStreamLength), 
		Round((str.Alt(0)+increasement.Alt(0))*EthalonStreamLength), 
		Round((str.Des(0)+increasement.Des(0))*EthalonStreamLength), 
	}}
	multiplier := str.Len(0)*EthalonStreamLength / newLen
	str.ScaleTo(Round(newLen))
	return multiplier
}
// complex improve:  TBD
// Ca, Cd, Cad/Cda 
// Ad, Ac, Adc/Acd 
// Da, Dc, Dac/Dca