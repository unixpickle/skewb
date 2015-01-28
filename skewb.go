package skewb

// A Corner is one of the eight corner pieces on the puzzle.
//
// A corner's orientation is determined by the direction the x sticker is
// facing. 0 means it's normal to the x axis, 1 the y axis, and 2 the z axis.
//
// A corner's index is the binary triple XYZ with the origin at the bottom back
// left corner of the cube.
type Corner struct {
	Piece       int8
	Orientation int8
}

// A Skewb represents the state of a Skewb.
type Skewb struct {
	Centers [6]int8
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

// TurnB performs a rotation of the B face which corresponds to the bottom back
// left side of the Skewb.
func (s *Skewb) TurnB(clock bool) {
	// TODO: this
}

// TurnL performs a rotation of the L face which corresponds to the bottom front
// left corner of the Skewb.
func (s *Skewb) TurnL(clock bool) {
	// TODO: this
}

// TurnR performs a rotation of the R face which corresponds to the bottom back
// right corner of the Skewb.
func (s *Skewb) TurnR(clock bool) {
	// Permute centers and corners.
	s.Centers[4], s.Centers[3], s.Centers[1] = s.Centers[1], s.Centers[4],
		s.Centers[3]
	s.Corners[3], s.Corners[0], s.Corners[5] = s.Corners[5], s.Corners[3],
		s.Corners[0]
	
	// Change corner orientations.
	for _, i := range []int{0, 1, 3, 5} {
		if s.Corners[i].Orientation == 0 {
			s.Corners[i].Orientation = 2
		} else if s.Corners[i].Orientation == 2 {
			s.Corners[i].Orientation = 1
		} else {
			s.Corners[i].Orientation = 0
		}
	}
}

// TurnU performs a rotation of the F face which corresponds to the top back
// left corner of the Skewb.
func (s *Skewb) TurnU(clock bool) {
	// TODO: this
}

// Solved returns true if the skewb is solved and oriented correctly.
func (s *Skewb) Solved() bool {
	// Check corners.
	for i := uint8(0); i < 8; i++ {
		if res.Corners[i].Piece != i || res.Corners[i].Orientation != 0 {
			return false
		}
	}
	
	// Check centers
	for i := uint8(0); i < 6; i++ {
		if res.Centers[i] != i {
			return false
		}
	}
	
	return true
}
