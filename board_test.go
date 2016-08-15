package mancala

import "testing"

func TestBoard_Finished(t *testing.T) {
	cases := []struct {
		board    Board
		expected bool
	}{
		{NewBoard(), false},
		{Board{0, 0, 0, 0, 0, 0, 24, 4, 4, 4, 4, 4, 4, 0}, true},
		{Board{4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 24}, true},
		{Board{0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 24}, true},
		{Board{0, 4, 0, 4, 0, 4, 12, 0, 0, 0, 0, 0, 0, 24}, true},
		{Board{0, 0, 0, 0, 0, 0, 24, 4, 0, 4, 0, 4, 0, 12}, true},
		{Board{0, 4, 0, 4, 0, 4, 12, 4, 0, 4, 0, 4, 0, 12}, false},
	}
	for _, c := range cases {
		output := c.board.Finished()
		if output != c.expected {
			t.Errorf("%#v.Finished() == %#v, expected %#v",
				c.board, output, c.expected)
		}
	}
}

func TestBoard_Winner(t *testing.T) {
	cases := []struct {
		board    Board
		expected Player
	}{
		{NewBoard(), -1},
		{Board{0, 0, 0, 0, 0, 0, 24, 4, 4, 4, 4, 4, 4, 0}, -1},
		{Board{4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 24}, -1},
		{Board{0, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 0, 24}, -1},
		{Board{0, 4, 0, 4, 0, 4, 12, 0, 0, 0, 0, 0, 0, 24}, -1},
		{Board{0, 0, 0, 0, 0, 0, 24, 4, 0, 4, 0, 4, 0, 12}, -1},
		{Board{0, 4, 0, 4, 0, 4, 12, 4, 0, 4, 0, 4, 0, 12}, -1},

		{Board{0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 0, 0, 0, 0}, South},
		{Board{0, 0, 0, 0, 0, 0, 32, 4, 4, 4, 0, 0, 0, 4}, South},
		{Board{4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 20}, South},

		{Board{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48}, North},
		{Board{4, 4, 4, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 32}, North},
		{Board{0, 0, 0, 0, 0, 0, 20, 4, 4, 4, 4, 4, 4, 4}, North},
	}
	for _, c := range cases {
		output := c.board.Winner()
		if output != c.expected {
			t.Errorf("%#v.Winner() == %#v, expected %#v",
				c.board, output, c.expected)
		}
	}
}
