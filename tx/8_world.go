package tx 

import (
	"rhymald/mag-epsilon/characters"
)

type Location struct {
	// to be moved to cache service redis or badger:
	// Cache map[string]*characters.Character 
	Grid struct {
		X map[int][]*characters.Character
		y map[int][]*characters.Character
		Z map[int][]*characters.Character
	}
	// Status struct {
	// 	Lobby   []*characters.Character // safe zones, no regen
	// 	Players []*characters.Character
	// 	Spawn []*characters.SpawnPoint // scratch, shape with entropy
	// 	Foes  []*characters.Character
	// }
	// Wells []*Stream
	// Drops []*DropPoint
}