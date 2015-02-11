package skewb

// COHeuristic stores a move count for each corner configuration.
type COHeuristic map[int]int

// MakeCOHeuristic uses breadth-first search to generate a COHeuristic.
func MakeCOHeuristic(start Skewb, moves []Move) COHeuristic {
	res := COHeuristic{}
	
	// Create the starting nodes
	nodes := []searchNode{searchNode{start, 0}}
	
	// Do the search
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		idx := encodeCO(&node.state)
		if _, ok := res[idx]; ok {
			continue
		}
		res[idx] = node.depth
		for _, move := range moves {
			newState := node.state
			newState.Move(move)
			nodes = append(nodes, searchNode{newState, node.depth + 1})
		}
	}
	
	return res
}

// MinMoves returns the minimum number of moves needed to orient the corners in
// some orientation.
func (c COHeuristic) MinMoves(s *Skewb) int {
	return c[encodeCO(s)]
}

func encodeCO(s *Skewb) int {
	res := 0
	mul := 1
	for i := 0; i < 8; i++ {
		res += int(s.Corners[i].Orientation) * mul
		mul *= 3
	}
	return res
}

type searchNode struct {
	state Skewb
	depth int
}
