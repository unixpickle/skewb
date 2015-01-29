package skewb

// A Corner is one of the eight corner pieces on the puzzle.
//
// A corner's orientation is determined by the direction the x sticker is
// facing. 0 means it's normal to the x axis, 1 the y axis, and 2 the z axis.
//
// A corner's index is the binary triple XYZ with the origin at the bottom back
// left corner of the cube.
type Corner struct {
	Piece       uint8
	Orientation uint8
}

// A Skewb represents the state of a Skewb.
type Skewb struct {
	Centers [6]uint8
	Corners [8]Corner
}

// NewSkewb returns a solved skewb.
func NewSkewb() *Skewb {
	var res Skewb

	// Create corners
	for i := uint8(0); i < 8; i++ {
		res.Corners[i].Piece = i
	}

	// Create centers
	for i := uint8(0); i < 6; i++ {
		res.Centers[i] = i
	}

	return &res
}

// Move applies a move to the Skewb.
func (s *Skewb) Move(m Move) {
	switch m.Face {
	case 'B':
		s.TurnB(m.Clock)
	case 'L':
		s.TurnL(m.Clock)
	case 'R':
		s.TurnR(m.Clock)
	case 'U':
		s.TurnU(m.Clock)
	}
}

// Solved returns true if the skewb is solved and oriented correctly.
func (s *Skewb) Solved() bool {
	// Check corners.
	for i := uint8(0); i < 8; i++ {
		if s.Corners[i].Piece != i || s.Corners[i].Orientation != 0 {
			return false
		}
	}

	// Check centers
	for i := uint8(0); i < 6; i++ {
		if s.Centers[i] != i {
			return false
		}
	}

	return true
}

// TurnB performs a rotation of the B face which corresponds to the bottom back
// left side of the Skewb.
func (s *Skewb) TurnB(clock bool) {
	// Permute centers and corners.
	if clock {
		s.Centers[5], s.Centers[1], s.Centers[3] = s.Centers[3], s.Centers[5],
			s.Centers[1]
		s.Corners[2], s.Corners[4], s.Corners[1] = s.Corners[1], s.Corners[2],
			s.Corners[4]
	} else {
		s.Centers[3], s.Centers[5], s.Centers[1] = s.Centers[5], s.Centers[1],
			s.Centers[3]
		s.Corners[1], s.Corners[2], s.Corners[4] = s.Corners[2], s.Corners[4],
			s.Corners[1]
	}

	// Change corner orientations.
	for _, i := range []int{0, 1, 2, 4} {
		if clock {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 1) % 3
		} else {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 2) % 3
		}
	}
}

// TurnL performs a rotation of the L face which corresponds to the bottom front
// left corner of the Skewb.
func (s *Skewb) TurnL(clock bool) {
	// Permute centers and corners.
	if clock {
		s.Centers[2], s.Centers[1], s.Centers[5] = s.Centers[5], s.Centers[2],
			s.Centers[1]
		s.Corners[6], s.Corners[5], s.Corners[0] = s.Corners[0], s.Corners[6],
			s.Corners[5]
	} else {
		s.Centers[5], s.Centers[2], s.Centers[1] = s.Centers[2], s.Centers[1],
			s.Centers[5]
		s.Corners[0], s.Corners[6], s.Corners[5] = s.Corners[6], s.Corners[5],
			s.Corners[0]
	}

	// Change corner orientations.
	for _, i := range []int{0, 4, 6, 5} {
		if clock {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 2) % 3
		} else {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 1) % 3
		}
	}
}

// TurnR performs a rotation of the R face which corresponds to the bottom back
// right corner of the Skewb.
func (s *Skewb) TurnR(clock bool) {
	// Permute centers and corners.
	if clock {
		s.Centers[4], s.Centers[3], s.Centers[1] = s.Centers[1], s.Centers[4],
			s.Centers[3]
		s.Corners[3], s.Corners[0], s.Corners[5] = s.Corners[5], s.Corners[3],
			s.Corners[0]
	} else {
		s.Centers[1], s.Centers[4], s.Centers[3] = s.Centers[4], s.Centers[3],
			s.Centers[1]
		s.Corners[5], s.Corners[3], s.Corners[0] = s.Corners[3], s.Corners[0],
			s.Corners[5]
	}

	// Change corner orientations.
	for _, i := range []int{0, 1, 3, 5} {
		if clock {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 2) % 3
		} else {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 1) % 3
		}
	}
}

// TurnU performs a rotation of the U face which corresponds to the top back
// left corner of the Skewb.
func (s *Skewb) TurnU(clock bool) {
	// Permute centers and corners.
	if clock {
		s.Centers[0], s.Centers[5], s.Centers[3] = s.Centers[3], s.Centers[0],
			s.Centers[5]
		s.Corners[6], s.Corners[0], s.Corners[3] = s.Corners[3], s.Corners[6],
			s.Corners[0]
	} else {
		s.Centers[3], s.Centers[0], s.Centers[5] = s.Centers[0], s.Centers[5],
			s.Centers[3]
		s.Corners[3], s.Corners[6], s.Corners[0] = s.Corners[6], s.Corners[0],
			s.Corners[3]
	}

	// Change corner orientation.
	for _, i := range []int{0, 2, 3, 6} {
		if clock {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 2) % 3
		} else {
			s.Corners[i].Orientation = (s.Corners[i].Orientation + 1) % 3
		}
	}
}
