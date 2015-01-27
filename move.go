package skewb

// A Move represents a move on the skewb.
// The side can be 0, 1, 2, or 3 corresponding to DFL, DFR, UFL, UFR.
type Move struct {
	Side  int
	Clock bool
}

// String returns my crappy notation for the move.
func (m Move) String() string {
	faceStr := "DFL"
	if m.Side == 1 {
		faceStr = "DFR"
	} else if m.Side == 2 {
		faceStr = "UFL"
	} else if m.Side == 3 {
		faceStr = "UFR"
	}
	if m.Clock {
		return faceStr
	} else {
		return faceStr + "'"
	}
}
