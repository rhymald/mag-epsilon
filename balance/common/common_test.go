package common

import "testing"
import "math"

func Test_BRandNewStream(t *testing.T){
	elems:=len(Elements)
		t.Logf("-------------------------------------------------")
		for x:=0; x<10; x++ {
		buffer := BRandNewStream(Elements[x%elems], 1)
		// for g:=0; g<(x+1)*(x+1)*(x+1); g++ { buffer.GrowAStream() }
		t.Logf("Stream of %s with length %.3f", buffer.Elem(), buffer.Len(0))
		t.Logf("  cre: %.3f, - alt %.3f, - des %.3f", buffer.Cre(0), buffer.Alt(0), buffer.Des(0))
		for d:=0; d<5; d++ {
			dot := buffer.EmitDot()
			if dot.Weight() != 0 {
				t.Logf("    - %.1f produces %s dot of %.0f weight", math.Log2(buffer.Len(1)), dot.Elem(), dot.Weight() )
			}
		}
		t.Logf("-------------------------------------------------")
	}
}