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

func Test_Actions(t *testing.T){
	pass, pass2, reset := fancy.Clr(1), fancy.Clr(1), fancy.Clr(0) 
	plus := float64(GrowStep) / EthalonStreamLength
	action := NewAction("0000-000000000=0-0000000", "fractal=Interruption", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	stream := BRandNewStream(Elements[0], MinEnthropy)
	for x:=0; x<0; x++ {
		dot := stream.EmitDot()
		action.Feed(stream, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Interrupted - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Interrupt("1234-123456789=1-1234567", [3]int{4,-9,1})
	t.Logf(" - %+v", *action)
	t.Logf("Interrupted: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == false { pass2 = fancy.Clr(8) } else { pass2 = fancy.Clr(2) }
	t.Logf("Interrupted - %sValid: %+v%s, %sSucceeded: %+v%s (false)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000=0-0000000", "fractal=Low", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		dot := stream.EmitDot()
		action.Feed(stream, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.9, 3, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Low energy: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s, Succeeded: %+v (any)", pass, action.Valid(), reset, action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000-0-0000000", "fractal=Weak", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		dot := stream.EmitDot()
		action.Feed(stream, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Weak - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.1, 12, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Weak streams: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	t.Logf("Weak streams - %sValid: %+v%s, Succeeded: %+v (any)", pass, action.Valid(), reset, action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000=0-0000000", "fractal=SelfCast", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		dot := stream.EmitDot()
		action.Feed(stream, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Self - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(1, 15, 15, [3]int{4,9,17}, [3]int{0,0,0})
	t.Logf(" - %+v", *action)
	t.Logf("Self: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("Self - %sValid: %+v%s, %sSucceeded: %+v%s (true)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000=0-0000000", "fractal=Success", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		dot := stream.EmitDot()
		action.Feed(stream, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("100%% success - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(1, 15, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("100%% success: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("100%% success - %sValid: %+v%s, %sSucceeded: %+v%s (true)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
}

func Test_EmitDot(t *testing.T){
	plus := float64(GrowStep) / EthalonStreamLength
	stream := BRandNewStream(Elements[0], MinEnthropy)
	t.Logf("%sScale changing:%s entropy direction changes, and grows", fancy.B, fancy.E[0])
	pass := ""
	for x:=0; x<32; x++ {
		dot := stream.EmitDot()
		if (stream.Des()+1) > (stream.Mean()+1) { pass = fancy.Clr(1) } else { pass = fancy.Clr(3) }
		t.Logf("%s | %9.3f mass | %9.3f weight | with %9.3f len | %s%+6.1f%%%s := %9.3f / %0.3f \t%+v %+v", dot.Elem(), DotWeightFromStreamLen(stream.Len()+1), dot.Weight(), stream.Len(), pass, 100*(stream.Des()+1)/(stream.Mean()+1)-100, fancy.E[0], (stream.Des()+1), (stream.Mean()+1), *dot, *stream)
		stream.Plus( (plus) )
	}
	stream = BRandNewStream(Elements[0], MinEnthropy)
	t.Logf("%sScale same:%s entropy direction stay, and grows", fancy.B, fancy.E[0])
	for x:=0; x<16; x++ {
		dot := stream.EmitDot()
		if (stream.Des()+1) > (stream.Mean()+1) { pass = fancy.Clr(1) } else { pass = fancy.Clr(3) }
		t.Logf("%s | %9.3f mass | %9.3f weight | with %9.3f len | %s%+6.1f%%%s := %9.3f / %0.3f \t%+v %+v", dot.Elem(), DotWeightFromStreamLen(stream.Len()+1), dot.Weight(), stream.Len(), pass, 100*(stream.Des()+1)/(stream.Mean()+1)-100, fancy.E[0], (stream.Des()+1), (stream.Mean()+1), *dot, *stream)
		stream.ScaleTo( 2*(stream.Len()) )
	}
}