package mancala

import "fmt"

type Player int

const (
	South Player = 0
	North Player = 1
)

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

type Seeds int

const (
	numberOfPlayers = 2
	pitsPerPlayer   = 6
	storesPerPlayer = 1
	boardSize       = numberOfPlayers * (pitsPerPlayer + storesPerPlayer)
)

type Board [boardSize]Seeds

type Turn struct {
	Board
	Player
	Pit
}

func NewBoard() Board {
	return Board{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
}

func (board *Board) collect(pit Pit) Seeds {
	seeds := board[pit]
	board[pit] = 0
	return seeds
}

func (p Player) Opponent() Player {
	return 1 - p
}

func (p Player) Pits() [pitsPerPlayer]Pit {
	switch p {
	case South:
		return SouthPits
	case North:
		return NorthPits
	default:
		return [pitsPerPlayer]Pit{}
	}
}

func (p Player) Store() Pit {
	switch p {
	case South:
		return SouthStore
	case North:
		return NorthStore
	default:
		return -1
	}
}

func (prev Pit) NextFor(player Player) (next Pit) {
	next = (prev + 1) % boardSize
	if next == player.Opponent().Store() {
		next = (next + 1) % boardSize
	}
	return
}

func (t Turn) Evaluate() (Board, bool) {
	board := t.Board
	pit := t.Pit
	for board[pit] > 1 && pit != t.Player.Store() {
		seeds := board.collect(pit)
		for seeds > 0 {
			pit = pit.NextFor(t.Player)
			seeds -= 1
			board[pit] += 1
		}
	}
	return board, pit == t.Player.Store()
}

func (board Board) EmptyFor(player Player) bool {
	for _, pit := range player.Pits() {
		if board[pit] > 0 {
			return false
		}
	}
	return true
}

func (board Board) Finished() bool {
	return board.EmptyFor(South) || board.EmptyFor(North)
}

func (board Board) String() string {
	return fmt.Sprintf(`
    %3d %3d %3d %3d %3d %3d
%3d                         %3d
    %3d %3d %3d %3d %3d %3d
`,
		board[NorthPitF],
		board[NorthPitE],
		board[NorthPitD],
		board[NorthPitC],
		board[NorthPitB],
		board[NorthPitA],
		board[NorthStore],
		board[SouthStore],
		board[SouthPitA],
		board[SouthPitB],
		board[SouthPitC],
		board[SouthPitD],
		board[SouthPitE],
		board[SouthPitF],
	)
}
