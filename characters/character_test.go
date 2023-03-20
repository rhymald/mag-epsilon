package characters

import (
	"math"
  "testing"
	"rhymald/mag-epsilon/fancy"
	"rhymald/mag-epsilon/common"
)

func Test_BRandNewStats(t *testing.T){
	physs:=len(common.Physical)
		t.Logf("-------------------------------------------------------------------------")
		for x:=0; x<10; x++ {
		buffer := BRandNewStats(common.Physical[x%physs])
		t.Logf("Generated player %s with body, %sand %d streams%s", buffer.GetID(), fancy.Clr(6-len(buffer.Streams)), len(buffer.Streams), fancy.Clr(0))
		t.Logf("  Body %s %.3f x %.3f x %.3f  | len %.3f", buffer.Body.Elem(), buffer.Body.Cre(0), buffer.Body.Alt(0), buffer.Body.Des(0), buffer.Body.Len(0))
		for i, each := range buffer.Streams {
			t.Logf("    - %d'%s %.3f x %.3f x %.3f | harm %.1f%% | len %.3f | %.3f dot", i+1, each.Elem(), each.Cre(0), each.Alt(0), each.Des(0), each.Harmony()*100, each.Len(0), math.Log2(each.Len(1)))
		}
		t.Logf("-------------------------------------------------------------------------")
	}
}

func Test_Character(t *testing.T){
	var char Character 
	char.Base = BRandNewStats(common.Physical[2])
	t.Logf("Character ID: %s", char.Base.GetID())
	char.Atts = char.Base.CalculaterAttributes(false)
	char.Cons = BrandNewLife()
	/// TBD
}