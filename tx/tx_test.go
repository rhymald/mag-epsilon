package tx

import (
	"testing"
	"rhymald/mag-epsilon/balance/common"
	"rhymald/mag-epsilon/fancy"
	"fmt"
)

func Test_Actions(t *testing.T){
	pass, pass2, reset := fancy.Clr(1), fancy.Clr(1), fancy.Clr(0) 
	plus := float64(common.GrowStep) / common.EthalonStreamLength
	action := NewAction("0000-000000000/0-0000000", "fractal=Interruption", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	stream := common.BRandNewStream(common.Elements[0], common.MinEnthropy)
	for x:=0; x<0; x++ {
		dot := stream.EmitDot()
		action.Feed(fmt.Sprintf("%d", common.Epoch()%3), dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %s: %+v", stream.Elem(), *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Interrupted - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Interrupt("1234-123456789/1-1234567", [3]int{4,-9,1})
	t.Logf(" - %+v", *action)
	t.Logf("Interrupted: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == false { pass2 = fancy.Clr(8) } else { pass2 = fancy.Clr(2) }
	t.Logf("Interrupted - %sValid: %+v%s, %sSucceeded: %+v%s (false)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000/0-0000000", "fractal=Low", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<9; x++ {
		dot := stream.EmitDot()
		action.Feed(fmt.Sprintf("%d", common.Epoch()%3), dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %s: %+v", stream.Elem(), *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.9, 3, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Low energy: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s, Succeeded: %+v (any)", pass, action.Valid(), reset, action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000/0-0000000", "fractal=Weak", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<9; x++ {
		dot := stream.EmitDot()
		action.Feed(fmt.Sprintf("%d", common.Epoch()%3), dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %s: %+v", stream.Elem(), *action)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Weak - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.1, 12, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Weak streams: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	t.Logf("Weak streams - %sValid: %+v%s, Succeeded: %+v (any)", pass, action.Valid(), reset, action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("0000-000000000/0-0000000", "fractal=SelfCast", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<9; x++ {
		dot := stream.EmitDot()
		action.Feed(fmt.Sprintf("%d", common.Epoch()%3), dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %s: %+v", stream.Elem(), *action)
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
	action = NewAction("0000-000000000/0-0000000", "fractal=Success", "withFlock=Default")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<9; x++ {
		dot := stream.EmitDot()
		action.Feed(fmt.Sprintf("%d", common.Epoch()%3), dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %s: %+v", stream.Elem(), *action)
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
