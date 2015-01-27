package skewb

type Skewb [30]int8

func (s *Skewb) TwistUFR() {
	s[2], s[22], s[12] = s[12], s[2], s[22] // check
	s[21], s[14], s[3] = s[3], s[21], s[14] // check
	s[23], s[11], s[1] = s[1], s[23], s[11] // check
	s[15], s[6], s[26] = s[26], s[15], s[6] // check
	s[4], s[20], s[11] = s[11], s[4], s[20] // check
}

func (s *Skewb) TwistUFL() {
	s[12], s[27], s[2] = s[2], s[12], s[27] // check
	s[4], s[13], s[25] = s[25], s[4], s[13] // check
	s[11], s[29], s[0] = s[0], s[11], s[29] // check
	s[20], s[5], s[16] = s[16], s[20], s[5] // check
	s[3], s[10], s[26] = s[26], s[3], s[10] // check
}

func (s *Skewb) TwistDFL() {
	s[12], s[7], s[27] = s[27], s[12], s[7] // check
	s[14], s[8], s[26] = s[26], s[14], s[8] // check
	s[23], s[19], s[3] = s[3], s[23], s[19] // check
	s[6], s[28], s[10] = s[10], s[6], s[28] // check
	s[13], s[5], s[29] = s[29], s[13], s[5] // check
}

func (s *Skewb) TwistDFR() {
	s[12], s[22], s[7] = s[7], s[12], s[22] // check
	s[20], s[9], s[13] = s[13], s[20], s[9] // check
	s[4], s[18], s[29] = s[29], s[4], s[18] // check
	s[11], s[24], s[5] = s[5], s[11], s[24] // check
	s[14], s[23], s[6] = s[6], s[14], s[23] // check
}
