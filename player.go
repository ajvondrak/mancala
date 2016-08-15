package mancala

type Player int

const (
	South Player = 0
	North Player = 1

	numberOfPlayers = 2
)

func (p Player) Opponent() Player {
	return 1 - p
}

func (p Player) Pits() [pitsPerPlayer]Pit {
	switch p {
	case South:
		return SouthPits
	case North:
		return NorthPits
	default:
		return [pitsPerPlayer]Pit{}
	}
}

func (p Player) Store() Pit {
	switch p {
	case South:
		return SouthStore
	case North:
		return NorthStore
	default:
		return -1
	}
}
