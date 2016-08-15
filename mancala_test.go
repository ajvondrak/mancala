package mancala

import "fmt"

func ExampleTurn_Evaluate() {
	turn := Turn{NewBoard(), South, SouthPitF}
	board, goAgain := turn.Evaluate()
	fmt.Printf("before:%safter:%sextra turn: %v", turn.Board, board, goAgain)
	// Output:
	// before:
	//       4   4   4   4   4   4
	//   0                           0
	//       4   4   4   4   4   4
	// after:
	//       5   5   5   0   5   5
	//   0                           2
	//       5   0   5   5   5   1
	// extra turn: true
}
