package mancala

import "testing"

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
