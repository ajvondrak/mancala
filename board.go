package mancala

import "fmt"

type Seeds int

const boardSize = numberOfPlayers * (pitsPerPlayer + storesPerPlayer)

type Board [boardSize]Seeds

func NewBoard() Board {
	return Board{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
}

func (board *Board) collect(pit Pit) Seeds {
	seeds := board[pit]
	board[pit] = 0
	return seeds
}

func (board Board) emptyFor(player Player) bool {
	for _, pit := range player.Pits() {
		if board[pit] > 0 {
			return false
		}
	}
	return true
}

func (board Board) Finished() bool {
	return board.emptyFor(South) || board.emptyFor(North)
}

func (board Board) scoreFor(player Player) Seeds {
	score := board[player.Store()]
	for _, leftover := range player.Pits() {
		score += board[leftover]
	}
	return score
}

func (board Board) Winner() Player {
	southScore := board.scoreFor(South)
	northScore := board.scoreFor(North)
	switch {
	case southScore > northScore:
		return South
	case northScore > southScore:
		return North
	default:
		return -1 // tie
	}
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
