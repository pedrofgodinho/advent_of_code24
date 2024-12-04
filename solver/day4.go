package solver

import (
	"fmt"
	"strings"
)

type Day4 struct {
	letters [][]byte
}

func (d *Day4) Parse(input string) (bool, error) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		d.letters = append(d.letters, []byte(line))
	}
	return true, nil
}

func (d *Day4) checkWordStartingOn(word []byte, row, col int) int {
	count := 0
	// Check east
	if col+len(word) <= len(d.letters[row]) {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row][col+i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check south
	if row+len(word) <= len(d.letters) {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row+i][col] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check west
	if col-len(word) >= -1 {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row][col-i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check north
	if row-len(word) >= -1 {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row-i][col] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check south-east
	if row+len(word) <= len(d.letters) && col+len(word) <= len(d.letters[row]) {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row+i][col+i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check south-west
	if row+len(word) <= len(d.letters) && col-len(word) >= -1 {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row+i][col-i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check north-west
	if row-len(word) >= -1 && col-len(word) >= -1 {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row-i][col-i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	// Check north-east
	if row-len(word) >= -1 && col+len(word) <= len(d.letters[row]) {
		match := true
		for i := 0; i < len(word); i++ {
			if d.letters[row-i][col+i] != word[i] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}

func (d *Day4) checkCrossOn(row, col int) bool {
	if row+2 >= len(d.letters) || col+2 >= len(d.letters[row]) {
		return false
	}

	return d.letters[row+1][col+1] == 'A' &&
		((d.letters[row][col] == 'M' && d.letters[row+2][col+2] == 'S') || (d.letters[row][col] == 'S' && d.letters[row+2][col+2] == 'M')) &&
		((d.letters[row][col+2] == 'M' && d.letters[row+2][col] == 'S') || (d.letters[row][col+2] == 'S' && d.letters[row+2][col] == 'M'))

}

func (d *Day4) Part1() (string, error) {
	word := []byte("XMAS")
	count := 0
	for row := 0; row < len(d.letters); row++ {
		for col := 0; col < len(d.letters[row]); col++ {
			count += d.checkWordStartingOn(word, row, col)
		}
	}
	return fmt.Sprintf("%d", count), nil
}

func (d *Day4) Part2() (string, error) {
	count := 0
	for row := 0; row < len(d.letters); row++ {
		for col := 0; col < len(d.letters[row]); col++ {
			if d.checkCrossOn(row, col) {
				count++
			}
		}
	}
	return fmt.Sprintf("%d", count), nil
}
