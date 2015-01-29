package skewb

import (
	"errors"
	"fmt"
	"strconv"
)

// CornerStickers maps the sticker colors to each corner.
var CornerStickers = []int{
	5, 1, 3,
	4, 1, 3,
	5, 0, 3,
	4, 0, 3,
	5, 1, 2,
	4, 1, 2,
	5, 0, 2,
	4, 0, 2,
}

// StickerCorners maps the sticker indices to each physical corner.
var StickerCorners = []int{
	28, 8, 19,
	24, 9, 18,
	25, 0, 16,
	21, 1, 15,
	29, 5, 13,
	23, 6, 14,
	26, 3, 10,
	20, 4, 11,
}

// CenterStickers maps each center piece to a sticker index.
var CenterStickers = []int{
	2, 7, 12, 17, 22, 27,
}

// ReadPuzzle uses standard input to read a puzzle from the user.
// This will return nil if the user entered an invalid puzzle.
func ReadPuzzle() (*Skewb, error) {
	var stickers [30]int
	for i := 0; i < 6; i++ {
		fmt.Printf("Enter face %d: ", i+1)
		var line string
		fmt.Scanln(&line)
		if len(line) != 5 {
			return nil, errors.New("Invalid length for face " +
				strconv.Itoa(i+1))
		}
		for j, ch := range line {
			m := map[rune]rune{'w': '1', 'y': '2', 'g': '3', 'b': '4',
				'r': '5', 'o': '6'}
			if x, ok := m[ch]; ok {
				ch = x
			}
			if ch < '1' || ch > '6' {
				return nil, errors.New("Invalid character: " + string(ch))
			}
			stickers[i*5+j] = int(ch - '1')
		}
	}
	return SkewbFromStickers(stickers)
}

// SkewbFromStickers turns an array containing the stickers of the Skewb into a
// a piece-by-piece Skewb.
func SkewbFromStickers(stickers [30]int) (*Skewb, error) {
	var result Skewb

	// Read and validate the centers
	var counts [6]int
	for i := 0; i < 6; i++ {
		center := stickers[CenterStickers[i]]
		result.Centers[i] = uint8(center)
		counts[center]++
	}
	for i := 0; i < 6; i++ {
		if counts[i] != 1 {
			return nil, errors.New("Unexpected number of center " +
				strconv.Itoa(i))
		}
	}

	// Read the corners
	for i := 0; i < 8; i++ {
		corner := [3]int{
			stickers[StickerCorners[i*3]],
			stickers[StickerCorners[i*3+1]],
			stickers[StickerCorners[i*3+2]],
		}
		found, piece, orientation := findCorner(corner)
		if !found {
			return nil, errors.New("Invalid corner: " +
				strconv.Itoa(corner[0]) + " " + strconv.Itoa(corner[1]) + " " +
				strconv.Itoa(corner[2]))
		}
		result.Corners[i].Piece = uint8(piece)
		result.Corners[i].Orientation = uint8(orientation)
	}

	return &result, nil
}

func findCorner(corner [3]int) (found bool, piece int, orientation int) {
	for piece = 0; piece < 8; piece++ {
		colors := CornerStickers[piece*3 : piece*3+3]
		if !setsEqual(colors, corner[:]) {
			continue
		}
		for orientation = 0; orientation < 3; orientation++ {
			if corner[orientation] == colors[0] {
				break
			}
		}
		return true, piece, orientation
	}
	return false, 0, 0
}

func setsEqual(s1 []int, s2 []int) bool {
	for _, x := range s1 {
		found := false
		for _, y := range s2 {
			if y == x {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
