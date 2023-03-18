package characters

import "testing"
import "rhymald/mag-epsilon/common"

func Test_BRandNewStats(t *testing.T){
	for x:=0; x<10; x++ {
		buffer := BRandNewStats((x+1)*10)
		t.Logf("Generated player %s with %d streams", buffer.GetID(), len(buffer.Streams))
		for i, each := range buffer.Streams {
			t.Logf("    #%d'%s %.3f x %.3f x %.3f     len %.3f | %.3f vol", i+1, each.Elem(), each.Cre(0), each.Alt(0), each.Des(0), each.Len(0), each.Vol(0))
		}
	}
}

func Test_Character(t *testing.T){
	var char Character 
	char.Base = BRandNewStats(common.Epoch()%2000+1)
	t.Logf("Character ID: %s", char.Base.GetID())
	char.Atts = char.Base.CalculaterAttributes(false)
	char.Cons = BrandNewLife()

}