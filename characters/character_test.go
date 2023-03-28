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
		t.Logf("  Body %s %.3f x %.3f x %.3f  | len %.3f", buffer.Body.Elem(), buffer.Body.Cre(), buffer.Body.Alt(), buffer.Body.Des(), buffer.Body.Len())
		buffer.BrandAStream(false)
		for i, each := range buffer.Streams {
			t.Logf("    - %d'%s %.3f x %.3f x %.3f | len %.3f | %.3f dot", i+1, each.Elem(), each.Cre(), each.Alt(), each.Des(), each.Len(), common.DotWeightFromStreamLen(each.Len()+1))
		}
		t.Logf("-------------------------------------------------------------------------")
	}
}

func Test_GainDot_BurnDot(t *testing.T){
	ups, brands, forks := 0, 0, 0
	var char Character 
	char.Base = BRandNewStats(common.Physical[common.Epoch()%5])
	for x:=0; x<ups; x++ { 
		char.Base.GrowAStream(false) 
		if common.Rand() < float64(forks)/float64(ups+1) { char.Base.SproutAStream(false) }
		if common.Rand() < float64(brands)/float64(ups+1) { char.Base.BrandAStream(false) }
	}
	t.Logf("%sCharacter ID:%s %s", fancy.B, fancy.E[0], char.Base.GetID())
	t.Logf("-------------------------------------------------------------------------")
	t.Logf("%sID:%s %+v", fancy.B, fancy.E[0], *&char.Base.ID)
	t.Logf("%sBody:%s %+v", fancy.B, fancy.E[0], *&char.Base.Body)
	streams := ""
	for _, each := range *&char.Base.Streams { streams = fmt.Sprintf("%s %+v", streams, *each) }
	t.Logf("%sStreams:%s %s", fancy.B, fancy.E[0], streams)
	t.Logf("-------------------------------------------------------------------------")
	char.Atts = char.Base.CalculaterAttributes()
	t.Logf("%sAttributes:%s %+v", fancy.B, fancy.E[0], *&char.Atts)
	t.Logf("-------------------------------------------------------------------------")
	streamCount := len(*&char.Base.Streams)
	char.Cons = BrandNewLife(streamCount)
	start, counter := common.Epoch(), 0
	t.Logf("%sFlocks:%s %+v", fancy.B, fancy.E[0], *(*&char.Cons.Flocks[0]))
	for x:=0; x<100; x++ {
		dots := ""
		for _, each := range *&char.Cons.Pool { dots = fmt.Sprintf("%s %d%s", dots, (*each)[each.Elem()], each.Elem())}
		t.Logf("%sHP:%s %5.1f%% | %sDots %+d:%s%s", fancy.B, fancy.E[0], float64(*&char.Cons.HP)/10, fancy.B, counter, fancy.E[0], dots)
		if x%2 == 0 { _ = char.Cons.BurnDot() }
		if x%3 == 0 { _ = char.Cons.BurnDot() }
		if x%4 == 0 { _ = char.Cons.BurnDot() }
		if x%6 == 0 { _ = char.Cons.BurnDot() }
		if x%9 == 0 { _ = char.Cons.BurnDot() }
		char.Cons.GainDotFrom(*&char.Base.Streams[x%streamCount]) ; common.Wait(128)
		counter++
		if common.Epoch()-start > 1024000000 {break}
	}
}