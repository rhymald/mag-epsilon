package common

import (
	"testing"
	"rhymald/mag-epsilon/fancy"
)

func Test_BRandNewStream(t *testing.T){
	elems:=len(Elements)
		t.Logf("-------------------------------------------------")
		for x:=0; x<10; x++ {
		buffer := BRandNewStream(Elements[x%elems], MinEnthropy)
		t.Logf("Stream of %s with length %.3f", buffer.Elem(), buffer.Len())
		t.Logf("  cre: %.3f, - alt %.3f, - des %.3f", buffer.Cre(), buffer.Alt(), buffer.Des())
		for d:=0; d<3; d++ {
			dot := buffer.EmitDot()
			if (*dot)[dot.Elem()] != 0 {
				t.Logf("    - %.3f produces %s dot of %.3f weight", DotWeightFromStreamLen(buffer.Len()+1), dot.Elem(), dot.Weight())
			}
		}
		t.Logf("-------------------------------------------------")
	}
}

func Test_EmitDot(t *testing.T){
	plus := float64(GrowStep) / EthalonStreamLength
	stream := BRandNewStream(Elements[0], MinEnthropy)
	t.Logf("%sScale changing:%s entropy direction changes, and grows", fancy.B, fancy.E[0])
	pass := ""
	for x:=0; x<32; x++ {
		Wait(32)
		dot := stream.EmitDot()
		if (stream.Des()+1) > (stream.Mean()+1) { pass = fancy.Clr(1) } else { pass = fancy.Clr(3) }
		t.Logf("%s | mass %9.3f %+5.1f%%  %0.3f  \tweight | with %9.3f len | %s%+6.1f%%%s := %9.3f / %0.3f \t%+v\t%+v", dot.Elem(), DotWeightFromStreamLen(stream.Len()+1), 100-(DotWeightFromStreamLen(stream.Len()+1))/dot.Weight()*100, dot.Weight(), stream.Len(), pass, 100*(stream.Des()+1)/(stream.Mean()+1)-100, fancy.E[0], (stream.Des()+1), (stream.Mean()+1), (*dot)[dot.Elem()], (*stream)[stream.Elem()])
		stream.Plus( (plus) )
	}
	stream = BRandNewStream(Elements[0], MinEnthropy)
	t.Logf("%sScale same:%s entropy direction stay, and grows", fancy.B, fancy.E[0])
	for x:=0; x<16; x++ {
		Wait(32)
		dot := stream.EmitDot()
		if (stream.Des()+1) > (stream.Mean()+1) { pass = fancy.Clr(1) } else { pass = fancy.Clr(3) }
		t.Logf("%s | mass %9.3f %+5.1f%%  %0.3f  \tweight | with %9.3f len | %s%+6.1f%%%s := %9.3f / %0.3f \t%+v\t%+v", dot.Elem(), DotWeightFromStreamLen(stream.Len()+1), 100-(DotWeightFromStreamLen(stream.Len()+1))/dot.Weight()*100, dot.Weight(), stream.Len(), pass, 100*(stream.Des()+1)/(stream.Mean()+1)-100, fancy.E[0], (stream.Des()+1), (stream.Mean()+1), (*dot)[dot.Elem()], (*stream)[stream.Elem()])
		stream.ScaleTo( 2*(stream.Len()) )
	}
}