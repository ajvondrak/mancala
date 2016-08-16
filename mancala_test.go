package mancala

import "fmt"

func ExampleTurn_Evaluate_moving_multiple_seeds() {
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

func ExampleTurn_Evaluate_moving_one_seed() {
	turn := Turn{
		Board{0, 1, 1, 0, 1, 0, 24, 1, 2, 1, 2, 0, 4, 11},
		South,
		SouthPitB,
	}
	board, goAgain := turn.Evaluate()
	fmt.Printf("before:%safter:%sextra turn: %v", turn.Board, board, goAgain)
	// Output:
	// before:
	//       4   0   2   1   2   1
	//  11                          24
	//       0   1   1   0   1   0
	// after:
	//       4   0   2   1   2   1
	//  11                          25
	//       0   0   0   1   0   1
	// extra turn: true
}

func ExampleDecision_Next() {
	d := Start
	fmt.Printf("\n%2d ---\n%s\n", 0, d.nextBoard)
	for i := 0; i < 50; i += 1 {
		d = d.Next(10)
		fmt.Printf("\n%2d ---%#v %#v\n%s\n",
			i, d.Turn.Player, d.Turn.Pit, d.nextBoard)
	}
	// Output: fuck
}
