package skewb

// Skewbs contains zero or more Skewb objects.
type Skewbs []Skewb

// Contains returns true if the list of skewbs contains a skewb equal to aSkewb.
func (s Skewbs) Contains(aSkewb *Skewb) bool {
	for i := 0; i < len(s); i++ {
		if SkewbsEqual(&s[i], aSkewb) {
			return true
		}
	}
	return false
}
