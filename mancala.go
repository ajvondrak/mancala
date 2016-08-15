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

type Seeds int

const numberOfPits = 14

type Board [numberOfPits]Seeds

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
	next = (prev + 1) % numberOfPits
	if next == player.Opponent().Store() {
		next = (next + 1) % numberOfPits
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
