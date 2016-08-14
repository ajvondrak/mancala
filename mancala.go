package main

import "fmt"

type Player int

type Board [14]int

func NewBoard() Board {
	return Board{4, 4, 4, 4, 4, 4, 0, 4, 4, 4, 4, 4, 4, 0}
}

func (board Board) String() string {
	return fmt.Sprintf(`
    %3d %3d %3d %3d %3d %3d
%3d                         %3d
    %3d %3d %3d %3d %3d %3d
`,
		board[12],
		board[11],
		board[10],
		board[9],
		board[8],
		board[7],
		board[13],
		board[6],
		board[0],
		board[1],
		board[2],
		board[3],
		board[4],
		board[5],
	)
}

func oppositeStore(player int) int {
	if player == 0 {
		return 13
	}
	return 6
}

func nextToSow(player int, previous int) int {
	next := (previous + 1) % 14
	if next == oppositeStore(player) {
		return (next + 1) % 14
	}
	return next
}

func (board *Board) Sow(player, pit int) {
	start := player*6 + pit
	seeds := board[start]
	board[start] = 0
	sown := start
	for seeds > 0 {
		sown = nextToSow(player, sown)
		seeds -= 1
		board[sown] += 1
	}
	// relay sowing
	if board[sown] > 1 {
		fmt.Printf("\n---\n")
		fmt.Printf(board.String())
		board.Sow(player, sown)
	}
}

func main() {
	board := NewBoard()
	fmt.Printf(board.String())
	board.Sow(0, 0)
	fmt.Printf("\n---\n")
	fmt.Printf(board.String())
}
