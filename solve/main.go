package main

import (
	"fmt"
	"github.com/unixpickle/skewb"
	"os"
)

var identity skewb.Skewb
var heuristic skewb.COHeuristic

func main() {
	// Input the puzzle.
	puzzle, err := skewb.ReadPuzzle()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	
	// Choose the identity with the same URF corner as the scramble.
	foundId := false
	for _, rot := range skewb.NewSkewb().AllRotations() {
		if rot.Corners[7].Piece == puzzle.Corners[7].Piece &&
			rot.Corners[7].Orientation == puzzle.Corners[7].Orientation {
			identity = rot
			foundId = true
			break
		}
	}
	if !foundId {
		fmt.Fprintln(os.Stderr, "Invalid URF corner.")
		os.Exit(1)
	}
	
	fmt.Println("Generating heuristic...")
	heuristic = skewb.MakeCOHeuristic(identity, skewb.AllMoves())
	for depth := 0; depth < 20; depth++ {
		fmt.Println("Exploring depth", depth, "...")
		solution := solve(puzzle, '_', depth)
		if solution != nil {
			fmt.Println("Got solution:", solution)
			return
		}
	}
}

func solve(s *skewb.Skewb, last rune, remaining int) []skewb.Move {
	if remaining == 0 {
		if !skewb.SkewbsEqual(s, &identity) {
			return nil
		} else {
			return []skewb.Move{}
		}
	} else {
		if heuristic.MinMoves(s) > remaining {
			return nil
		}
	}

	for _, face := range "BLRU" {
		if face == last {
			continue
		}
		for i := 0; i < 2; i++ {
			m := skewb.Move{face, i == 0}
			newS := *s
			newS.Move(m)
			if solution := solve(&newS, face, remaining-1); solution != nil {
				return append([]skewb.Move{m}, solution...)
			}
		}
	}

	return nil
}
