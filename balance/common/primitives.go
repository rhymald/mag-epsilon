package common

import (
	"math" // for dot only
	"math/rand"
	"time"
)

type Dot map[string]int 
type Stream map[string][3]int
type Element struct {
	Aka string
	//   ^ Name
	Eats, EatenBy, MadeOf, Next []string 
	// ^    ^        ^ Native: native - penetrates easier, lucky effect
  // ^    ^ Makes stream unstable, but powerful: fuel - easiest penetration
} // ^ Makes strem weak, but chills: chill - harder penetration
// else: no - regular penetration


    ////////////// 
   // ELEMENTS // 
	//////////////
var Elements []string = []string{"◌ ", "🌪 ", "🔥", "🪨", "🧊", "🌑", "🩸", "🎶", "☀️ "} 
var ElemList map[string]Element = map[string]Element{
	// Epoch 1: before Echaen's reserches: humans and feathers learn destruction
	Elements[0]: Element{Aka: "Common", EatenBy: []string{Elements[6]}, Next: Elements[0:3]},
	Elements[1]: Element{Aka: "Air",    Eats: []string{Elements[3]}, EatenBy: []string{Elements[2], Elements[7]}},
	Elements[2]: Element{Aka: "Fire",   Eats: []string{Elements[1]}, EatenBy: []string{Elements[4], Elements[6]}},
	// Epoch 2: +alt +cre - less focus on element, more schools, +fractals series
	Elements[3]: Element{Aka: "Earth",  Eats: []string{Elements[4]}, EatenBy: []string{Elements[1], Elements[7]}},
	Elements[4]: Element{Aka: "Water",  Eats: []string{Elements[2]}, EatenBy: []string{Elements[3], Elements[6]}},
	// Epoch 3: enslaved wild shadows learn to live with humans
	Elements[5]: Element{Aka: "Void",   Eats: Elements[:], MadeOf: []string{Elements[0]}},
	// Epoch 4: rhyxxix break out - ethernal shadows release
	Elements[6]: Element{Aka: "Mallom", Eats: []string{Elements[2], Elements[4]}, EatenBy: []string{Elements[5], Elements[2]}, MadeOf: []string{Elements[2], Elements[4]}},
	// Epoch 5: mind controllers on 4, annihilators from 1
	Elements[7]: Element{Aka: "Noise",  Eats: []string{Elements[1], Elements[3]}, EatenBy: []string{Elements[1], Elements[8]}, MadeOf: []string{Elements[1], Elements[3]}},
	Elements[8]: Element{Aka: "Resonance", Eats: Elements[1:5], EatenBy: Elements, MadeOf: []string{Elements[2]}},
}
var Physical []string = []string{"◌ ", "🌱", "🪵", "🪨", "🛡 "} // none, flesh/plant (punchD), wood/shell (stingA), stone (chopDa), armored
var PhysList map[string]Element = map[string]Element{
	// Sandy, ooze, flaiming or whirling spirits
	Physical[0]: Element{Aka: "Spirit"},
	// Living creatures:
	Physical[1]: Element{Aka: "Flesh/Plant"}, 
	Physical[2]: Element{Aka: "Wood/Shell"}, 
	// Summoned and elementals:
	Physical[3]: Element{Aka: "Hard",   MadeOf: []string{Physical[1]}}, 
	// Plate armor and mechanisms
	Physical[4]: Element{Aka: "Forged", MadeOf: []string{Physical[1], Physical[2]}}, 
}

    //////////
   // DOTS //
  //////////
func DotWeightFromStreamLen(a float64) float64 { return  4/3 * math.Pi * math.Pow(Log(a/2),3) }
func (str *Stream) EmitDot() *Dot { return &Dot{ str.Elem(): ChancedRound( 1000*DotWeightFromStreamLen( str.Len(1) ) ) } }
func (dot *Dot) Weight() float64 { buf := *dot ; return float64(buf[dot.Elem()])/1000 }
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
func (str *Stream) Mean() float64 { return 3/(1/str.Cre(0)+1/str.Alt(0)+1/str.Des(0)) }
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
func (str *Stream) Attune() {
	picker := rand.New(rand.NewSource(time.Now().UnixNano())).Intn( len(ElemList[str.Elem()].Next) )
	newElem := ElemList[str.Elem()].Next[picker]
	buffer := Stream{ newElem : (*str)[str.Elem()] }
	*str = buffer
}
// complex improve:  TBD
// Ca, Cd, Cad/Cda 
// Ad, Ac, Adc/Acd 
// Da, Dc, Dac/Dca