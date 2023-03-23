package characters

import (
  "testing"
	"rhymald/mag-epsilon/fancy"
	"rhymald/mag-epsilon/balance/common"
	"fmt"
)

func Test_BRandNewStats(t *testing.T){
	physs := len(common.Physical)
	t.Logf("-------------------------------------------------------------------------")
	for x:=0; x<10; x++ {
		buffer := BRandNewStats(common.Physical[x%physs])
		t.Logf("Generated player %s with body, %sand %d streams%s", buffer.GetID(), fancy.Clr(6-len(buffer.Streams)), len(buffer.Streams), fancy.Clr(0))
		t.Logf("  Body %s %.3f x %.3f x %.3f  | len %.3f", buffer.Body.Elem(), buffer.Body.Cre(0), buffer.Body.Alt(0), buffer.Body.Des(0), buffer.Body.Len(0))
		for i, each := range buffer.Streams {
			t.Logf("    - %d'%s %.3f x %.3f x %.3f | harm %.1f%% | len %.3f | %.3f dot", i+1, each.Elem(), each.Cre(0), each.Alt(0), each.Des(0), each.Harmony()*100, each.Len(0), common.DotWeightFromStreamLen(each.Len(1)))
		}
		t.Logf("-------------------------------------------------------------------------")
	}
}

func Test_Character(t *testing.T){
	var char Character 
	char.Base = BRandNewStats(common.Physical[2])
	t.Logf("%sCharacter ID:%s %s", fancy.B, fancy.E[0], char.Base.GetID())
	t.Logf("-------------------------------------------------------------------------")
	t.Logf("%sID:%s %+v", fancy.B, fancy.E[0], *&char.Base.ID)
	t.Logf("%sBody:%s %+v", fancy.B, fancy.E[0], *&char.Base.Body)
	streams := ""
	for _, each := range *&char.Base.Streams { streams = fmt.Sprintf("%s %+v", streams, *each) }
	t.Logf("%sStreams:%s %s", fancy.B, fancy.E[0], streams)
	t.Logf("-------------------------------------------------------------------------")
	char.Atts = char.Base.CalculaterAttributes(false)
	t.Logf("%sAttributes:%s %+v", fancy.B, fancy.E[0], *&char.Atts)
	t.Logf("-------------------------------------------------------------------------")
	char.Cons = BrandNewLife()
	streamCount := len(*&char.Base.Streams)
	for x:=0; x<12; x++ {
		dots := ""
		for _, each := range *&char.Cons.Pool { dots = fmt.Sprintf("%s %+v", dots, *each) }
		t.Logf("%sHP:%s %5.1f%% | %sDots:%s%s", fancy.B, fancy.E[0], float64(*&char.Cons.HP)/10, fancy.B, fancy.E[0], dots)
		if x%2 == 0 { _, _ = char.Cons.BurnDot() }
		if x%3 == 0 { _, _ = char.Cons.BurnDot() }
		if x%4 == 0 { _, _ = char.Cons.BurnDot() }
		if x%6 == 0 { _, _ = char.Cons.BurnDot() }
		if x%9 == 0 { _, _ = char.Cons.BurnDot() }
		char.Cons.GetDotFrom(*&char.Base.Streams[x%streamCount], char.Atts)
		if float64(len(*&char.Cons.Pool)) >= *&char.Atts.Poolsize+1 {break}
	}
}