package skewb

// A Move represents a move on the skewb.
// The face can be 'B', 'L', 'R', 'U'.
type Move struct {
	Face  rune
	Clock bool
}

// AllMoves returns all the available moves.
func AllMoves() []Move {
	res := make([]Move, 0, 8)
	for _, face := range "BLRU" {
		for i := 0; i < 2; i++ {
			res = append(res, Move{face, i == 0})
		}
	}
	return res
}

// String returns my crappy notation for the move.
func (m Move) String() string {
	if m.Clock {
		return string(m.Face)
	} else {
		return string(m.Face) + "'"
	}
}
