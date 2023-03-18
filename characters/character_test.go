package characters

import "testing"
import "rhymald/mag-epsilon/fancy"

func Test_BRandNewStats(t *testing.T){
	for x:=0; x<10; x++ {
		buffer := BRandNewStats()
		t.Logf("Generated player %s %swith %d streams%s", buffer.GetID(), fancy.Clr(6-len(buffer.Streams)), len(buffer.Streams), fancy.Clr(0))
		for i, each := range buffer.Streams {
			t.Logf("    #%d'%s %.3f x %.3f x %.3f     len %.3f | %.3f vol", i+1, each.Elem(), each.Cre(0), each.Alt(0), each.Des(0), each.Len(0), each.Vol(0))
		}
	}
}

func Test_Character(t *testing.T){
	var char Character 
	char.Base = BRandNewStats()
	t.Logf("Character ID: %s", char.Base.GetID())
	char.Atts = char.Base.CalculaterAttributes(false)
	char.Cons = BrandNewLife()

}