package mancala

import (
	"fmt"
	"testing"
)

func TestPlayer_Opponent(t *testing.T) {
	cases := []struct {
		player   Player
		expected Player
	}{
		{North, South},
		{South, North},
	}
	for _, c := range cases {
		output := c.player.Opponent()
		if output != c.expected {
			t.Errorf("%#v.Opponent() == %#v, expected %#v",
				c.player, output, c.expected)
		}
	}
}

func TestPlayer_Store(t *testing.T) {
	cases := []struct {
		player   Player
		expected Pit
	}{
		{South, SouthStore},
		{North, NorthStore},
		{0xdeadbeef, -1},
	}
	for _, c := range cases {
		output := c.player.Store()
		if output != c.expected {
			t.Errorf("%#v.Store() == %#v, expected %#v",
				c.player, output, c.expected)
		}
	}
}

func TestPit_NextFor(t *testing.T) {
	cases := []struct {
		pit      Pit
		player   Player
		expected Pit
	}{
		{SouthPitA, South, SouthPitB},
		{SouthPitB, South, SouthPitC},
		{SouthPitC, South, SouthPitD},
		{SouthPitD, South, SouthPitE},
		{SouthPitE, South, SouthPitF},
		{SouthPitF, South, SouthStore},
		{SouthStore, South, NorthPitA},
		{NorthPitA, South, NorthPitB},
		{NorthPitB, South, NorthPitC},
		{NorthPitC, South, NorthPitD},
		{NorthPitD, South, NorthPitE},
		{NorthPitE, South, NorthPitF},
		{NorthPitF, South, SouthPitA},
		{NorthStore, South, SouthPitA},

		{SouthPitA, North, SouthPitB},
		{SouthPitB, North, SouthPitC},
		{SouthPitC, North, SouthPitD},
		{SouthPitD, North, SouthPitE},
		{SouthPitE, North, SouthPitF},
		{SouthPitF, North, NorthPitA},
		{SouthStore, North, NorthPitA},
		{NorthPitA, North, NorthPitB},
		{NorthPitB, North, NorthPitC},
		{NorthPitC, North, NorthPitD},
		{NorthPitD, North, NorthPitE},
		{NorthPitE, North, NorthPitF},
		{NorthPitF, North, NorthStore},
		{NorthStore, North, SouthPitA},
	}
	for _, c := range cases {
		output := c.pit.NextFor(c.player)
		if output != c.expected {
			t.Errorf("%#v.NextFor(%#v) == %#v, expected %#v",
				c.pit, c.player, output, c.expected)
		}
	}
}

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
