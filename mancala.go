package mancala

type Turn struct {
	Board
	Player
	Pit
}

func (t Turn) Evaluate() (Board, bool) {
	board := t.Board
	pit := t.Pit
	for {
		seeds := board.collect(pit)
		for seeds > 0 {
			pit = pit.NextFor(t.Player)
			seeds -= 1
			board[pit] += 1
		}
		if board[pit] <= 1 || pit == t.Player.Store() {
			break
		}
	}
	return board, pit == t.Player.Store()
}

type Decision struct {
	Turn
	nextBoard  Board
	nextPlayer Player
}

var Start = Decision{nextBoard: NewBoard(), nextPlayer: South}

func (d Decision) ValueToMaximize() Seeds {
	return d.nextBoard.scoreFor(South)
}

func (d Decision) Terminal() bool {
	return d.nextBoard.Finished() ||
		d.nextBoard[SouthStore] >= 25 ||
		d.nextBoard[NorthStore] >= 25
}

func NewDecision(turn Turn) Decision {
	nextBoard, extraTurn := turn.Evaluate()
	if extraTurn {
		return Decision{turn, nextBoard, turn.Player}
	}
	return Decision{turn, nextBoard, turn.Player.Opponent()}
}

func (d Decision) Continuations() []Decision {
	continuations := make([]Decision, 0, pitsPerPlayer)
	for _, pit := range d.nextPlayer.Pits() {
		if d.nextBoard[pit] > 0 {
			turn := Turn{d.nextBoard, d.nextPlayer, pit}
			continuations = append(continuations, NewDecision(turn))
		}
	}
	return continuations
}

func (d Decision) Next(depth int) Decision {
	if depth == 0 || d.Terminal() {
		return d
	}

	choices := d.Continuations()

	switch d.nextPlayer {
	case South:
		var bestValue Seeds = -1
		var bestChoice Decision

		for _, c := range choices {
			value := c.Next(depth - 1).ValueToMaximize()
			if value > bestValue {
				bestValue = value
				bestChoice = c
			}
		}

		return bestChoice
	case North:
		var bestValue Seeds = 49
		var bestChoice Decision

		for _, c := range choices {
			value := c.Next(depth - 1).ValueToMaximize()
			if value < bestValue {
				bestValue = value
				bestChoice = c
			}
		}

		return bestChoice
	default:
		panic("fuck")
	}
}
