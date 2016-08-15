package mancala

type Pit int

const (
	SouthPitA Pit = iota
	SouthPitB
	SouthPitC
	SouthPitD
	SouthPitE
	SouthPitF
	SouthStore

	NorthPitA
	NorthPitB
	NorthPitC
	NorthPitD
	NorthPitE
	NorthPitF
	NorthStore
)

const (
	pitsPerPlayer   = 6
	storesPerPlayer = 1
)

var SouthPits = [...]Pit{
	SouthPitA,
	SouthPitB,
	SouthPitC,
	SouthPitD,
	SouthPitE,
	SouthPitF,
}

var NorthPits = [...]Pit{
	NorthPitA,
	NorthPitB,
	NorthPitC,
	NorthPitD,
	NorthPitE,
	NorthPitF,
}

func (prev Pit) NextFor(player Player) (next Pit) {
	next = (prev + 1) % boardSize
	if next == player.Opponent().Store() {
		next = (next + 1) % boardSize
	}
	return
}
