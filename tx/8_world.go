package tx 

import (
	"rhymald/mag-epsilon/balance/common"
	"errors"
	"rhymald/mag-epsilon/characters"
	"sync"
)

type Location struct {
	Queue struct {
		Actions map[int][]*Action
		Effects map[int][]*Effect
		sync.Mutex
	}
	// to be moved to cache service redis or badger:
	Cache struct {
		Lobby map[string]*characters.Character 
		Alive map[string]*characters.Character 
		NPC map[string]*characters.Character 
		sync.Mutex
	}
	Grid struct {
		X map[int][]*characters.Character
		Y map[int][]*characters.Character
		Z map[int][]*characters.Character
		sync.Mutex
	}
}


// NEW
func NewLocation() *Location {
	var buffer Location
	buffer.Cache.Lobby = make(map[string]*characters.Character) 
	buffer.Cache.Alive = make(map[string]*characters.Character) 
	buffer.Cache.NPC = make(map[string]*characters.Character) 
	buffer.Grid.X = make(map[int][]*characters.Character)
	buffer.Grid.Y = make(map[int][]*characters.Character)
	buffer.Grid.Z = make(map[int][]*characters.Character)
	// buffer.Queue.Actions = make(map[int][]*Action)
	buffer.Queue.Effects = make(map[int][]*Effect)
	return &buffer
}


// READ 
func (loc *Location) HasPlayerInCacheLobby(pl *characters.Character) string {
	loc.Cache.Lock() ; inLobby := (*loc).Cache.Lobby[(*pl).Base.GetID()[:14]] 
	inLobbyOld := (*loc).Cache.Lobby[(*pl).Base.GetID()[:14]+"-old"] ; loc.Cache.Unlock()
	if inLobby == nil && inLobbyOld == nil { return "FALSE" }
	err := "Error: unexpected"
	if inLobby != nil && inLobbyOld == nil { if (*inLobby).Base.GetID() == (*pl).Base.GetID() { return "TRUE: new" } else { err = "Error: base full id != requested id" } }
	if inLobby == nil && inLobbyOld != nil { 
		if (*inLobbyOld).Base.GetID()[:14] == (*pl).Base.GetID()[:14] { 
			if (*inLobbyOld).Base.GetID() != (*pl).Base.GetID() { 
				return "TRUE: old" 
			} else { 
				err = "Error: current == old"
			} 
		} else { 
			err = "Error: base born id != requested id"
		} 
	}
	if inLobby != nil && inLobbyOld != nil { 
		if (*inLobby).Base.GetID() == (*pl).Base.GetID() {
			if (*inLobbyOld).Base.GetID()[:14] == (*pl).Base.GetID()[:14] {
				if (*inLobbyOld).Base.GetID() != (*pl).Base.GetID() { 
					return "TRUE: both" 
				} else { 
					err = "Error: current == old"
				} 
			} else { 
				err = "Error: current born id != old"
			}
		} else {
			err = "Error: current base id != new"
		}}
	return err
}

func (loc *Location) HasPlayerInCacheAlive(pl *characters.Character) bool {
	isInAlive, isInNPC := false, false
	loc.Cache.Lock() ; isAlive := (*loc).Cache.Alive[(*pl).Base.GetID()] 
	isNPC := (*loc).Cache.NPC[(*pl).Base.GetID()] ; loc.Cache.Unlock()
	if isAlive != nil { if (*isAlive).Base.GetID() == (*pl).Base.GetID() { isInAlive = true } }
	if isNPC != nil { if (*isNPC).Base.GetID() == (*pl).Base.GetID() { isInNPC = true } }
	if ( isInAlive && !isInNPC ) || ( !isInAlive && isInNPC ) { return true }
	return false
}

func (loc *Location) SeekPlayerInGrid(pl *characters.Character) ([]int, []int, []int) {
	var x, y, z []int
	loc.Cache.Lock()
	for _, each := range (*loc).Grid.X { for _, char := range each { if (*char).Base.GetID()[:14] == (*pl).Base.GetID()[:14] { x = append(x, char.Where()[0]) }}}
	for _, each := range (*loc).Grid.Y { for _, char := range each { if (*char).Base.GetID()[:14] == (*pl).Base.GetID()[:14] { y = append(y, char.Where()[1]) }}}
	for _, each := range (*loc).Grid.Z { for _, char := range each { if (*char).Base.GetID()[:14] == (*pl).Base.GetID()[:14] { z = append(z, char.Where()[2]) }}}
	loc.Cache.Unlock()
	if len(x) == len(y) && len(y) == len(z) { return x, y, z }
	return []int{}, []int{}, []int{}
}
func (loc *Location) HasPlayerInGridAtXYZ(pl *characters.Character, xyz [3]int) (float64, float64) {
	ingrid, incoords := -1.0, -1.0
	for _, each := range loc.GetCharListFromArea(pl.Where(), 2) { if each == pl { ingrid = common.Between( pl.Where(), each.Where() ) ; break }}
	for _, each := range loc.GetCharListFromArea(xyz, 2) { if each == pl { incoords = common.Between( pl.Where(), each.Where() ) ; break }}
	return ingrid, incoords
}

func (loc *Location) GetCharListFromArea(XYZ [3]int, radius int) []*characters.Character {
	var buffer [][]*characters.Character 
	var sorted []*characters.Character
	loc.Cache.Lock()
	for x:=XYZ[0] ; x<XYZ[0]+radius ; x++ { if len((*loc).Grid.X[x]) != 0 { buffer = append(buffer, (*loc).Grid.X[x]) }}
	for y:=XYZ[1] ; y<XYZ[1]+radius ; y++ { if len((*loc).Grid.Y[y]) != 0 { buffer = append(buffer, (*loc).Grid.Y[y]) }}
	for z:=XYZ[2] ; z<XYZ[2]+radius ; z++ { if len((*loc).Grid.Z[z]) != 0 { buffer = append(buffer, (*loc).Grid.Z[z]) }}
	loc.Cache.Unlock()
	for _, bunch := range buffer { for _, each := range bunch {
		inList := false 
		for i,_ := range sorted { if sorted[i] == each { inList = true }}
		if !inList { sorted = append(sorted, each) }
	}}
	return sorted
}


// MOD
func (loc *Location) PutCharToLobbyCache(pl *characters.Character) error {
	plid := (*pl).Base.GetID()[:14]
	if (*pl).Base.IsNPC() { return errors.New("Is NPC! Only players can get in lobby.") }
	check := loc.HasPlayerInCacheLobby(pl)
	if check[:5] == "Error" { return errors.New("Invalid thing in cahce found: "+check) }//; fix it cleanup }
	if check == "TRUE: old" { return errors.New("It should not be in old only") }
	loc.Cache.Lock() 
	list := (*loc).Cache.Lobby 
	if check[:5] == "FALSE" { list[plid] = pl }
	if check == "TRUE: new" { list[plid+"-old"], list[plid] = list[plid], pl }
	if check == "TRUE: both" { list[plid+"-extra"], list[plid+"-old"], list[plid] = list[plid+"-old"], list[plid], pl }
	(*loc).Cache.Lobby = list
	loc.Cache.Unlock()
	return nil
}

func (loc *Location) PutCharToLifeCache(pl *characters.Character) error {
	check := loc.HasPlayerInCacheLobby(pl)
	if (*pl).Base.IsNPC() {

	} else {
		if check != "TRUE: new" && check != "TRUE: both" { return errors.New("No such player in Lobby: "+check) }
		// 
	}
	return nil
}

func (loc *Location) PutCharToGrid(pl *characters.Character) error {
	if loc.HasPlayerInCacheAlive(pl) { return errors.New("Actual char is absent in any life Cache.") }
	return nil
}

// RemoveChar
// LogoutChar
// MoveChar