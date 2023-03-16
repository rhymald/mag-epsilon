package characters

import "testing"

func Test_BRandNewStats(t *testing.T){
	for x:=0; x<10; x++ {
		buffer := BRandNewStats((x+1)*10, (x%2==1))
		t.Logf("%v | %s with %s %.3f x %.3f x %.3f", buffer.ID.NPC, buffer.GetID(), buffer.Stats.Elem(), buffer.Stats.Cre(), buffer.Stats.Alt(), buffer.Stats.Des())
	}
}