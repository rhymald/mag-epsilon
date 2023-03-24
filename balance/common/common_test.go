package common

import "testing"
// import "math"

func Test_BRandNewStream(t *testing.T){
	elems:=len(Elements)
		t.Logf("-------------------------------------------------")
		for x:=0; x<10; x++ {
		buffer := BRandNewStream(Elements[x%elems], MinEnthropy)
		// for g:=0; g<(x+1)*(x+1)*(x+1); g++ { buffer.GrowAStream() }
		t.Logf("Stream of %s with length %.3f", buffer.Elem(), buffer.Len(0))
		t.Logf("  cre: %.3f, - alt %.3f, - des %.3f", buffer.Cre(0), buffer.Alt(0), buffer.Des(0))
		for d:=0; d<3; d++ {
			dot := buffer.EmitDot()
			if (*dot)[dot.Elem()] != 0 {
				t.Logf("    - %.3f produces %s dot of %.3f weight", DotWeightFromStreamLen(buffer.Len(1)), dot.Elem(), dot.Weight())
			}
		}
		t.Logf("-------------------------------------------------")
	}
}

func Test_Actions(t *testing.T){
	plus := 128
	action := NewAction("Interruption")
	t.Logf("Creating - %+v", *action)
	stream := BRandNewStream(Elements[0], MinEnthropy)
	for x:=0; x<0; x++ {
		dot := stream.EmitDot()
		action.Feed(x, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	t.Logf("Interrupted - Valid: %+v (false)", action.Valid())
	action.Interrupt("1234-123456789-1-1234567", [3]int{4,-9,1})
	t.Logf(" - %+v", *action)
	t.Logf("Interrupted: %+v", *&action.Result)
	t.Logf("Interrupted - Valid: %+v, Succeeded: %+v (false)", action.Valid(), action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("Low")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<1; x++ {
		dot := stream.EmitDot()
		action.Feed(x, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	t.Logf("Low energy - Valid: %+v (false)", action.Valid())
	action.Finish(0.9, 3, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Low energy: %+v", *&action.Result)
	t.Logf("Low energy - Valid: %+v, Succeeded: %+v (any)", action.Valid(), action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("Weak")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<4; x++ {
		dot := stream.EmitDot()
		action.Feed(x, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	t.Logf("Weak streams - Valid: %+v (false)", action.Valid())
	action.Finish(0.1, 12, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("Weak streams: %+v", *&action.Result)
	t.Logf("Weak streams - Valid: %+v, Succeeded: %+v (any)", action.Valid(), action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("SelfCast")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<7; x++ {
		dot := stream.EmitDot()
		action.Feed(x, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	t.Logf("Self - Valid: %+v (false)", action.Valid())
	action.Finish(1, 15, 15, [3]int{4,9,17}, [3]int{0,0,0})
	t.Logf(" - %+v", *action)
	t.Logf("Self: %+v", *&action.Result)
	t.Logf("Self - Valid: %+v, Succeeded: %+v (true)", action.Valid(), action.Succeeded())
	t.Logf("-------------------------------------------------")
	action = NewAction("Success")
	t.Logf("Creating - %+v", *action)
	for x:=0; x<9; x++ {
		dot := stream.EmitDot()
		action.Feed(x, dot)
		stream.Attune()
		stream.Plus(plus)
		t.Logf(" - %+v", *action)
  }
	t.Logf("100%% success - Valid: %+v (false)", action.Valid())
	action.Finish(1, 15, 15, [3]int{4,9,0}, [3]int{0,1,7})
	t.Logf(" - %+v", *action)
	t.Logf("100%% success: %+v", *&action.Result)
	t.Logf("100%% success - Valid: %+v, Succeeded: %+v (true)", action.Valid(), action.Succeeded())
}

func Test_Eitdot(t *testing.T){
	plus := 1024
	stream := BRandNewStream(Elements[0], MinEnthropy)
	for x:=0; x<10; x++ {
		dot := stream.EmitDot()
		t.Logf("produced %s: %.3f / %.3f weight - with %+v", dot.Elem(), DotWeightFromStreamLen(stream.Len(1)), dot.Weight(), *stream)
		// stream.Attune()
		stream.ScaleTo(2*int( float64(plus) * stream.Len(0) ))
	}
}