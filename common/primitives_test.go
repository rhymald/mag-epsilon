package common

import "testing"

func Test_BRangNewStream(t *testing.T){
	elems:=len(Elements)
	for x:=0; x<10; x++ {
		buffer := BRandNewStream((x+1)*10, Elements[x%elems])
		t.Logf("Stream of %s with - cre: %.3f, - alt %.3f, - des %.3f", buffer.Elem(), buffer.Cre(), buffer.Alt(), buffer.Des())
	}
}