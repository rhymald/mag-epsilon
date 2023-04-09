package characters

import (
	// "sync"
	"rhymald/mag-epsilon/balance/common"
	// "fmt"
)

type SpawnPoint struct {
	ID string
	// Message []string
	XYZ [3]int
	Radius int 
	Behavior string // [player] or [aware/curious/apathy, pred/herbivore, aggressive/calm/coward/stationary, lone/pack/swarm, ...]
	Concurrent []*SpawnPoint // same with colliding areas or login points
	// for npc only:
	Lifecycle int 
	Size common.Stream
	Attune common.Stream
	// Skills +Extra
}

var (
	Behaviours []string = []string{ "Player", // for non-npc
		"Aware", "Curious", "Apathy", // for player nearby
		"Predator", "Herbivore", // tbd: on hunt or defend
		"Aggressive", "Calm", "Coward", "Stationary", // battle movement pattern
		"Lone", "Pack", "Swarm", "Hive", // social: defence or consumation for victory
	}
	LoginPoints []*SpawnPoint = WelcomeGateWays("4")
	Spawn_Dummies []*SpawnPoint = Spawn_TrainingsDummies("4")
) 

func WelcomeGateWays(world string) []*SpawnPoint {
	var buffer []*SpawnPoint
	if world == "4" {
		ashtown :=  SpawnPoint{ ID: "Humans' Ashtown", XYZ: [3]int{7000,0,0}, Behavior: "Player", Radius: 500 }
		feathers := SpawnPoint{ ID: "Feathers' village", XYZ: [3]int{9000,0,0}, Behavior: "Player", Radius: 400 }
		swamp :=    SpawnPoint{ ID: "Swamp slums", XYZ: [3]int{4000,0,0}, Behavior: "Player", Radius:1350 }
		farland :=  SpawnPoint{ ID: "Tribe", XYZ: [3]int{1000,0,0}, Behavior: "Player", Radius: 200 }
		ashtown.Concurrent = []*SpawnPoint{ &feathers, &swamp, &farland }
		feathers.Concurrent = []*SpawnPoint{ &ashtown, &swamp, &farland }
		swamp.Concurrent = []*SpawnPoint{ &feathers, &ashtown, &farland }
		farland.Concurrent = []*SpawnPoint{ &feathers, &swamp, &ashtown }
		buffer = append(buffer, &ashtown)
		buffer = append(buffer, &feathers)
		buffer = append(buffer, &swamp)
		buffer = append(buffer, &farland)
		return buffer
	}
	// for _, each := range buffer { fmt.Println(*each) }
	return buffer
}

func Spawn_TrainingsDummies(world string) []*SpawnPoint {
	var buffer []*SpawnPoint
	if world == "4" {
		var coords map[int][3]int = map[int][3]int{ 
			100: [3]int{7000,-750,0}, 
			120: [3]int{9000, 500,0}, 
			 70: [3]int{4000,1500,0}, 
			200: [3]int{1000,-350,0},
		} // ^ radius:coords
		scratch := SpawnPoint{ 
			ID: "Training dummy",
			Behavior: "Calm|Apathy|Stationary", 
			Lifecycle: 1000 * 60 * 5, // ms
			Size:   common.Stream{ common.Physical[2]+"|"+common.Physical[4]: [3]int{0, 0, 0} }, // x to y to z
			Attune: common.Stream{ common.Elements[0]+"|"+common.Elements[3]+"|"+common.Elements[1]: [3]int{0, 0, 0} }, // x + y + z = 100.0%
		}
		for rad, xyz := range coords {
			plus := scratch
			plus.Radius, plus.XYZ = rad, xyz
			buffer = append(buffer, &plus)
		} 
	}
	return buffer
}