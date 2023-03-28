package common

import (
	"math" // for dot only
	// "math/rand"
	// "time"
)

type Stream map[string][3]int

const EthalonStreamLength float64 = 1024
const BaseStreamLength float64 = 512
const GrowStep float64 = 256
const MinEnthropy int = 4


// CREATE
func BRandNewStream(elem string, length int) *Stream {
	leng := GrowStep
	if elem == Elements[0] { leng = BaseStreamLength }
	enthropy := 1/float64(length+1)/float64(length+1)
	c, a, d := (1+Rand()-Rand())*enthropy, (1+Rand()-Rand())*enthropy, (1+Rand()-Rand())*enthropy
	for step:=0; step<length-1; step++ { c, a, d = c+(1+Rand()-Rand())*enthropy, a+(1+Rand()-Rand())*enthropy, d+(1+Rand()-Rand())*enthropy }
	modifier := 1 / Vector(c, a, d)
	return &Stream{ elem: [3]int{ CeilRound(c*leng*modifier), CeilRound(a*leng*modifier), CeilRound(d*leng*modifier) }}
} 


// READ
func (str *Stream) Elem() string { for elem, _ := range *str { return elem } ; return "ERR" }
func (str *Stream) Cre() float64 { return float64((*str)[str.Elem()][0])/EthalonStreamLength }
func (str *Stream) Alt() float64 { return float64((*str)[str.Elem()][1])/EthalonStreamLength }
func (str *Stream) Des() float64 { return float64((*str)[str.Elem()][2])/EthalonStreamLength }

func (str *Stream) Mean() float64 { return 3/(1/str.Cre()+1/str.Alt()+1/str.Des()) }
func (str *Stream) Len() float64 { return Vector(str.Cre(),str.Alt(),str.Des()) }


// EDIT
func (str *Stream) RandShapeAs(cc, aa, dd int) {
	keepLen := str.Len() * EthalonStreamLength
	entroc, entroa, entrod := 1/float64(cc+1)/float64(cc+1), 1/float64(aa+1)/float64(aa+1), 1/float64(dd+1)/float64(dd+1)
	c, a, d := (1+Rand()-Rand())*entroc, (1+Rand()-Rand())*entroa, (1+Rand()-Rand())*entrod 
	for step:=0; step<cc-1; step++ { c += entroc*(1+Rand()-Rand()) }
	for step:=0; step<aa-1; step++ { a += entroa*(1+Rand()-Rand()) }
	for step:=0; step<dd-1; step++ { d += entrod*(1+Rand()-Rand()) }
	modifier := 1 / Vector(c, a, d)
	*str = Stream{ str.Elem(): [3]int{ Round(c*keepLen*modifier), Round(a*keepLen*modifier), Round(d*keepLen*modifier) }}
}

// NEED to get fixed from int
func (str *Stream) ScaleTo(newlen float64) {
	multiplier := math.Round(newlen*EthalonStreamLength/GrowStep)/EthalonStreamLength*GrowStep / str.Len()
	c, a, d := str.Cre()*EthalonStreamLength, str.Alt()*EthalonStreamLength, str.Des()*EthalonStreamLength
	*str = Stream{ str.Elem(): [3]int{ CeilRound(c*multiplier), CeilRound(a*multiplier), CeilRound(d*multiplier) }}
}

// NEED to get fixed from int
func (str *Stream) Plus(ll float64) float64 {
	increasement := BRandNewStream(str.Elem(), MinEnthropy)
	increasement.ScaleTo( ll * float64(EthalonStreamLength) )
	newLen := math.Round((str.Len()+ll)*EthalonStreamLength/GrowStep)/EthalonStreamLength*GrowStep
	*str = Stream{ str.Elem(): [3]int{ 
		Round((str.Cre()+increasement.Cre())*EthalonStreamLength), 
		Round((str.Alt()+increasement.Alt())*EthalonStreamLength), 
		Round((str.Des()+increasement.Des())*EthalonStreamLength), 
	}}
	multiplier := str.Len()*EthalonStreamLength / newLen
	str.ScaleTo(newLen)
	return multiplier
}

func (str *Stream) Attune() {
	if len(ElemList[str.Elem()].Next) == 0 { return }
	picker := Epoch() % len(ElemList[str.Elem()].Next) 
	newElem := ElemList[str.Elem()].Next[picker]
	buffer := Stream{ newElem : (*str)[str.Elem()] }
	*str = buffer
}

// complex improve:  TBD
// Ca, Cd, Cad/Cda 
// Ad, Ac, Adc/Acd 
// Da, Dc, Dac/Dca