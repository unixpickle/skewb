package skewb

// A Skewb represents the state of a Skewb.
type Skewb [30]int8

// NewSkewb returns a solved skewb.
func NewSkewb() *Skewb {
	var res Skewb
	for i := int8(0); i < 30; i++ {
		res[i] = (i/5) + 1
	}
	return &res
}

// ClockDFL does a DFL move.
func (s *Skewb) ClockDFL() {
	s[12], s[7], s[27] = s[27], s[12], s[7] // check
	s[14], s[8], s[26] = s[26], s[14], s[8] // check
	s[23], s[19], s[3] = s[3], s[23], s[19] // check
	s[6], s[28], s[10] = s[10], s[6], s[28] // check
	s[13], s[5], s[29] = s[29], s[13], s[5] // check
}

// ClockDFR does a DFR move.
func (s *Skewb) ClockDFR() {
	s[12], s[22], s[7] = s[7], s[12], s[22] // check
	s[20], s[9], s[13] = s[13], s[20], s[9] // check
	s[4], s[18], s[29] = s[29], s[4], s[18] // check
	s[11], s[24], s[5] = s[5], s[11], s[24] // check
	s[14], s[23], s[6] = s[6], s[14], s[23] // check
}

// ClockUFL does a UFL move.
func (s *Skewb) ClockUFL() {
	s[12], s[27], s[2] = s[2], s[12], s[27] // check
	s[4], s[13], s[25] = s[25], s[4], s[13] // check
	s[11], s[29], s[0] = s[0], s[11], s[29] // check
	s[20], s[5], s[16] = s[16], s[20], s[5] // check
	s[3], s[10], s[26] = s[26], s[3], s[10] // check
}

// ClockUFR does a UFR move.
func (s *Skewb) ClockUFR() {
	s[2], s[22], s[12] = s[12], s[2], s[22] // check
	s[21], s[14], s[3] = s[3], s[21], s[14] // check
	s[23], s[10], s[1] = s[1], s[23], s[10] // check
	s[15], s[6], s[26] = s[26], s[15], s[6] // check
	s[4], s[20], s[11] = s[11], s[4], s[20] // check
}

// Move performs a move.
func (s *Skewb) Move(m Move) {
	count := 1
	if !m.Clock {
		count = 2
	}
	for i := 0; i < count; i++ {
		switch m.Side {
		case 0:
			s.ClockDFL()
		case 1:
			s.ClockDFR()
		case 2:
			s.ClockUFL()
		case 3:
			s.ClockUFR()
		}
	}
}

// Solved returns true if the skewb is solved in any orientation.
func (s *Skewb) Solved() bool {
	for i := 0; i < 6; i++ {
		color := s[i*5]
		for j := i*5+1; j < (i+1)*5; j++ {
			if s[j] != color {
				return false
			}
		}
	}
	return true
}
