package skewb

// A Move represents a move on the skewb.
// The face can be 'B', 'L', 'R', 'U'.
type Move struct {
	Face  rune
	Clock bool
}

// String returns my crappy notation for the move.
func (m Move) String() string {
	if m.Clock {
		return string(m.Face)
	} else {
		return string(m.Face) + "'"
	}
}
