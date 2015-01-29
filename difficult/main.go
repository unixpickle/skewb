package main

import (
	"fmt"
	"github.com/unixpickle/skewb"
)

func main() {
	counts := map[int64]int{}

	nodes := make([]searchNode, 1)
	nodes[0] = searchNode{*skewb.NewSkewb(), 0, ""}
	moves := skewb.AllMoves()
	lastDepth := -1
	for len(nodes) > 0 {
		n := nodes[0]
		nodes = nodes[1:]

		idx := hashState(&n.state)
		if _, ok := counts[idx]; ok {
			continue
		}

		if n.depth > lastDepth {
			fmt.Println("Up to depth", n.depth)
			lastDepth = n.depth
		}

		if n.depth == 11 {
			fmt.Println("Found node:", n.moves)
		}
		counts[idx] = n.depth
		for _, move := range moves {
			newState := n.state
			newState.Move(move)
			nodes = append(nodes, searchNode{newState, n.depth + 1,
				n.moves + " " + move.String()})
		}
	}
}

type searchNode struct {
	state skewb.Skewb
	depth int
	moves string
}

// hashCenters returns a number between 0 and 6^5
func hashCenters(s *skewb.Skewb) int64 {
	var res int64
	var scaler int64 = 1
	for i := 0; i < 5; i++ {
		res += scaler * int64(s.Centers[i])
		scaler *= 6
	}
	return res
}

// hashCO returns a number between 0 and 3^7
func hashCO(s *skewb.Skewb) int64 {
	var res int64
	var scaler int64 = 1
	for i := 0; i < 7; i++ {
		res += scaler * int64(s.Corners[i].Orientation)
		scaler *= 3
	}
	return res
}

// hashCP returns a number between 0 and 8^7
func hashCP(s *skewb.Skewb) int64 {
	var res int64
	var scaler int64 = 1
	for i := 0; i < 7; i++ {
		res += scaler * int64(s.Corners[i].Piece)
		scaler *= 8
	}
	return res
}

func hashState(s *skewb.Skewb) int64 {
	return hashCP(s) + hashCO(s)*2097152 + hashCenters(s)*4586471424
}
