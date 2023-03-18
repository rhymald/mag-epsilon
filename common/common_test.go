package common

import "testing"

func Test_BRandNewStream(t *testing.T){
	elems:=len(Elements)
	for x:=0; x<10; x++ {
		buffer := BRandNewStream(Elements[x%elems], x+1)
		t.Logf("Stream of %s with volume %.3f, length %.3f", buffer.Elem(), buffer.Vol(0), buffer.Len(0))
		t.Logf("       cre: %.3f, - alt %.3f, - des %.3f", buffer.Cre(0), buffer.Alt(0), buffer.Des(0))
		for d:=0; d<5; d++ {
			dot := buffer.EmitDot()
			t.Logf("       produces %s dot of %.0f weight", dot.Elem(), dot.Weight() )
		}
	}
}