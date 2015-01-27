package main

import (
	"fmt"
	"github.com/unixpickle/skewb"
)

func main() {
	puzzle := skewb.ReadPuzzle()
	if puzzle == nil {
		return
	}
	for depth := 0; depth < 20; depth++ {
		fmt.Println("Exploring depth", depth, "...")
		solution := solve(puzzle, -1, depth)
		if solution != nil {
			fmt.Println("Got solution:", solution)
			return
		}
	}
}

func solve(s *skewb.Skewb, last int, remaining int) []skewb.Move {
	if remaining == 0 {
		if !s.Solved() {
			return nil
		} else {
			return []skewb.Move{}
		}
	}
	
	for face := 0; face < 4; face++ {
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
