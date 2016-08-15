package mancala

import "testing"

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

func TestPlayer_Pits(t *testing.T) {
	cases := []struct {
		player   Player
		expected [pitsPerPlayer]Pit
	}{
		{South, SouthPits},
		{North, NorthPits},
		{0xdeadbeef, [...]Pit{0, 0, 0, 0, 0, 0}},
	}
	for _, c := range cases {
		output := c.player.Pits()
		if output != c.expected {
			t.Errorf("%#v.Pits() == %#v, expected %#v",
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
