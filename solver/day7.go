package solver

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Op int

const (
	Add    Op = iota
	Mult   Op = iota
	Concat Op = iota
)

type Line struct {
	target   int
	elements []int
}

type Day7 struct {
	lines []Line
}

func (d *Day7) Parse(input string) (bool, error) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return false, fmt.Errorf("expected 2 parts, got %d", len(parts))
		}

		target, ok := strconv.Atoi(parts[0])
		if ok != nil {
			return false, fmt.Errorf("expected target to be int, got %d", target)
		}

		var elements []int

		for _, element := range strings.Split(parts[1], " ") {
			e, ok := strconv.Atoi(element)
			if ok != nil {
				return false, fmt.Errorf("expected element to be int, got %v", element)
			}
			elements = append(elements, e)
		}

		d.lines = append(d.lines, Line{target, elements})
	}

	return true, nil
}

func GenerateCombinations[T any](values *[]T, n int) <-chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)

		combination := make([]T, n)
		var generate func(k int)
		generate = func(k int) {
			if k == n {
				// Create a copy only when a full combination is ready
				newCombination := make([]T, n)
				copy(newCombination, combination)
				ch <- newCombination
				return
			}

			for _, value := range *values {
				combination[k] = value // Modify the existing slice in place
				generate(k + 1)
			}
		}
		generate(0)
	}()

	return ch
}

func (d *Day7) Part1() (string, error) {
	allowedOps := []Op{Add, Mult}
	sum := 0
	for _, line := range d.lines {
		for combination := range GenerateCombinations(&allowedOps, len(line.elements)-1) {
			total := line.elements[0]
			for i := 1; i < len(line.elements); i++ {
				switch combination[i-1] {
				case Add:
					total += line.elements[i]
				case Mult:
					total *= line.elements[i]
				}
				if total > line.target {
					break
				}
			}
			if total == line.target {
				sum += total
				break
			}
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func concatenateInts(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("inputs must be non-negative")
	}

	// Handle the case where b is 0 separately.
	if b == 0 {
		if a > math.MaxInt/10 {
			return 0, fmt.Errorf("overflow: result exceeds maximum int value")
		}
		return a * 10, nil
	}

	multiplier := 1
	tempB := b

	// Optimized loop to find the multiplier
	for tempB > 0 {
		if math.MaxInt/multiplier < 10 {
			return 0, fmt.Errorf("overflow: result exceeds maximum int value")
		}
		multiplier *= 10
		tempB /= 10
	}

	if math.MaxInt-b < a*multiplier {
		return 0, fmt.Errorf("overflow: result exceeds maximum int value")
	}

	return a*multiplier + b, nil
}

func (d *Day7) Part2() (string, error) {
	allowedOps := []Op{Add, Mult, Concat}
	sum := 0
	for _, line := range d.lines {
		for combination := range GenerateCombinations(&allowedOps, len(line.elements)-1) {
			total := line.elements[0]
			for i := 1; i < len(line.elements); i++ {
				switch combination[i-1] {
				case Add:
					total += line.elements[i]
				case Mult:
					total *= line.elements[i]
				case Concat:
					total, _ = concatenateInts(total, line.elements[i])
				}
				if total > line.target {
					break
				}
			}
			if total == line.target {
				sum += total
				break
			}
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
