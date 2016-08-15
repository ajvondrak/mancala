package mancala

type Turn struct {
	Board
	Player
	Pit
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
