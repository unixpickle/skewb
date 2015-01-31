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

// SkewbsEqual compares two Skewbs directly.
func SkewbsEqual(s1 *Skewb, s2 *Skewb) bool {
	for i := 0; i < 8; i++ {
		if s1.Corners[i].Piece != s2.Corners[i].Piece ||
			s1.Corners[i].Orientation != s2.Corners[i].Orientation {
			return false
		}
	}
	for i := 0; i < 6; i++ {
		if s1.Centers[i] != s2.Centers[i] {
			return false
		}
	}
	return true
}

// AllRotations returns all 24 rotations of this Skewb.
func (s *Skewb) AllRotations() []Skewb {
	result := make([]Skewb, 0, 24)
	
	puzzle := *s
	
	// All the rotations except those with the top or bottom face in front.
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result = append(result, puzzle)
			puzzle.RotateZ()
		}
		puzzle.RotateY()
	}
	
	// Use the bottom face as the front and go through all four permutations.
	puzzle.RotateX()
	for i := 0; i < 4; i++ {
		result = append(result, puzzle)
		puzzle.RotateZ()
	}
	
	// Use the top face as the front and go through all four permutations
	puzzle.RotateX()
	puzzle.RotateX()
	for i := 0; i < 4; i++ {
		result = append(result, puzzle)
		puzzle.RotateZ()
	}
	
	return result
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

// RotateX performs a WCA "x" rotation to the Skewb.
// This move goes along what is the R face on the Rubik's cube.
func (s *Skewb) RotateX() {
	// Permute the centers.
	s.Centers[0], s.Centers[3], s.Centers[1], s.Centers[2] = s.Centers[2], 
		s.Centers[0], s.Centers[3], s.Centers[1]
	
	// Permute the corners.
	s.Corners[2], s.Corners[0], s.Corners[4], s.Corners[6] = s.Corners[6],
		s.Corners[2], s.Corners[0], s.Corners[4]
	s.Corners[3], s.Corners[1], s.Corners[5], s.Corners[7] = s.Corners[7],
		s.Corners[3], s.Corners[1], s.Corners[5]
	
	// Swap the y and z orientations.
	for i := 0; i < 8; i++ {
		if s.Corners[i].Orientation == 1 {
			s.Corners[i].Orientation = 2
		} else if s.Corners[i].Orientation == 2 {
			s.Corners[i].Orientation = 1
		}
	}
}

// RotateY performs a WCA "y" rotation to the Skewb.
// This move goes along what is the U face on the Rubik's cube.
func (s *Skewb) RotateY() {
	// Permute the centers.
	s.Centers[2], s.Centers[5], s.Centers[3], s.Centers[4] = s.Centers[4],
		s.Centers[2], s.Centers[5], s.Centers[3]
	
	// Permute the corners.
	s.Corners[2], s.Corners[3], s.Corners[7], s.Corners[6] = s.Corners[6],
		s.Corners[2], s.Corners[3], s.Corners[7]
	s.Corners[0], s.Corners[1], s.Corners[5], s.Corners[4] = s.Corners[4],
		s.Corners[0], s.Corners[1], s.Corners[5]
	
	// Swap the x and z orientations.
	for i := 0; i < 8; i++ {
		if s.Corners[i].Orientation == 0 {
			s.Corners[i].Orientation = 2
		} else if s.Corners[i].Orientation == 2 {
			s.Corners[i].Orientation = 0
		}
	}
}

// RotateZ performs a WCA "z" rotation to the Skewb.
// This move goes along what is the F face on the Rubik's cube.
func (s *Skewb) RotateZ() {
	// Permute the centers.
	s.Centers[0], s.Centers[4], s.Centers[1], s.Centers[5] = s.Centers[5],
		s.Centers[0], s.Centers[4], s.Centers[1]
	
	// Permute the corners.
	s.Corners[2], s.Corners[3], s.Corners[1], s.Corners[0] = s.Corners[0],
		s.Corners[2], s.Corners[3], s.Corners[1]
	s.Corners[6], s.Corners[7], s.Corners[5], s.Corners[4] = s.Corners[4],
		s.Corners[6], s.Corners[7], s.Corners[5]
	
	// Swap x and y orientations
	for i := 0; i < 8; i++ {
		if s.Corners[i].Orientation == 0 {
			s.Corners[i].Orientation = 1
		} else if s.Corners[i].Orientation == 1 {
			s.Corners[i].Orientation = 0
		}
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
