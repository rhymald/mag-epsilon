package tx

import (
	"testing"
	"rhymald/mag-epsilon/balance/common"
	"rhymald/mag-epsilon/characters"
	"rhymald/mag-epsilon/fancy"
	// "fmt"
)

func Test_Actions(t *testing.T){
	pass, pass2, reset := fancy.Clr(1), fancy.Clr(1), fancy.Clr(0) 
	stats := characters.BRandNewStats(common.Physical[1])
	action := NewAction("Interrupted", stats)
	t.Logf("Creating - %+v", *action)
	for x:=0; x<0; x++ {
		for y:=0; y<len((*stats).Streams); y++ {
			dot := (*stats).Streams[y].EmitDot()
			action.Feed(y+1, dot) ; common.Wait(5)
			t.Logf(" - %+v", *action)
		}
		stats.BrandAStream(false)
		stats.GrowAStream(false)
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
	action = NewAction("LowEnergy", stats)
	t.Logf("Creating - %+v", *action)
	for x:=0; x<1; x++ {
		for y:=0; y<len((*stats).Streams); y++ {
			dot := (*stats).Streams[y].EmitDot()
			action.Feed(y+1, dot) ; common.Wait(5)
			t.Logf(" - %+v", *action)
		}
		stats.BrandAStream(false)
		stats.GrowAStream(false)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.9, 3, 10, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Low energy: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("Low energy - %sValid: %+v%s, %sSucceeded: %+v%s (any)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("Weak", stats)
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		for y:=0; y<len((*stats).Streams); y++ {
			dot := (*stats).Streams[y].EmitDot()
			action.Feed(y+1, dot) ; common.Wait(5)
			t.Logf(" - %+v", *action)
		}
		stats.BrandAStream(false)
		stats.GrowAStream(false)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Weak - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(0.1, 10, 10, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Weak streams: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("Weak streams - %sValid: %+v%s, %sSucceeded: %+v%s (any)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("SelfCast", stats)
	t.Logf("Creating - %+v", *action)
	for x:=0; x<3; x++ {
		for y:=0; y<len((*stats).Streams); y++ {
			dot := (*stats).Streams[y].EmitDot()
			action.Feed(y+1, dot) ; common.Wait(5)
			t.Logf(" - %+v", *action)
		}
		stats.BrandAStream(false)
		stats.GrowAStream(false)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("Self - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(1, 10, 10, [3]int{4,9,17}, [3]int{0,0,0})
	t.Logf(" - %+v", *action)
	t.Logf("Self: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("Self - %sValid: %+v%s, %sSucceeded: %+v%s (any)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
	t.Logf("-------------------------------------------------")
	action = NewAction("Successful", stats)
	t.Logf("Creating - %+v", *action)
	for x:=0; x<2; x++ {
		for y:=0; y<len((*stats).Streams); y++ {
			dot := (*stats).Streams[y].EmitDot()
			action.Feed(y+1, dot) ; common.Wait(5)
			t.Logf(" - %+v", *action)
		}
		stats.BrandAStream(false)
		stats.GrowAStream(false)
  }
	if action.Valid() == false { pass = fancy.Clr(8) } else { pass = fancy.Clr(2) }
	t.Logf("100%% success - %sValid: %+v%s (false)", pass, action.Valid(), reset)
	action.Finish(1, 10, 10, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("100%% success: %+v", *&action.Result)
	if action.Valid() == true { pass = fancy.Clr(1) } else { pass = fancy.Clr(2) }
	if action.Succeeded() == true { pass2 = fancy.Clr(1) } else { pass2 = fancy.Clr(2) }
	t.Logf("100%% success - %sValid: %+v%s, %sSucceeded: %+v%s (any)", pass, action.Valid(), reset, pass2, action.Succeeded(), reset)
}
