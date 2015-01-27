package skewb

import (
	"fmt"
)

func ReadPuzzle() *Skewb {
	var puzzle Skewb
	for i := 0; i < 6; i++ {
		fmt.Printf("Enter face %d: ", i+1)
		var line string
		fmt.Scanln(&line)
		if len(line) != 5 {
			fmt.Println("You typed it wrong. You don't get a second chance.")
			return nil
		}
		for j, ch := range line {
			m := map[rune]rune{'w': '1', 'y': '2', 'g': '3', 'b': '4',
				'r': '5', 'o': '6'}
			if x, ok := m[ch]; ok {
				ch = x
			}
			if ch < '1' || ch > '6' {
				fmt.Println("Yup,", ch, "was wrong.")
				return nil
			}
			puzzle[i*5+j] = int8(ch-'1') + 1
		}
	}
	return &puzzle
}
