package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {
	diffsPerReport [][]int
}

func (d *Day2) Parse(input string) (bool, error) {
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)
		intFields := make([]int, len(fields))
		for i, field := range fields {
			newField, err := strconv.Atoi(field)
			if err != nil {
				return false, err
			}
			intFields[i] = newField
		}
		d.diffsPerReport = append(d.diffsPerReport, reportToDiffs(intFields))
	}
	return true, nil
}

func reportToDiffs(report []int) []int {
	diffs := make([]int, len(report)-1)
	for i := 0; i < len(report)-1; i++ {
		diffs[i] = report[i+1] - report[i]
	}
	return diffs
}

func All[T any](arr []T, pred func(T) bool) bool {
	for _, elem := range arr {
		if !pred(elem) {
			return false
		}
	}
	return true
}

func FindAndCount[T any](arr []T, pred func(T) bool) (int, int) {
	count := 0
	first := -1
	for i, elem := range arr {
		if pred(elem) {
			if first == -1 {
				first = i
			}
			count++
		}
	}
	return first, count
}

func AllExcept[T any](arr []T, except int, pred func(T) bool) bool {
	for i, elem := range arr {
		if i == except {
			continue
		}
		if !pred(elem) {
			return false
		}
	}
	return true
}

func (d *Day2) Part1() (string, error) {
	sum := 0
	for _, diffs := range d.diffsPerReport {
		firstSign := diffs[0] > 0
		good := All(diffs, func(diff int) bool {
			return diff > 0 == firstSign && diff != 0 && diff < 4 && diff > -4
		})
		if good {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func (d *Day2) Part2() (string, error) {
	sum := 0
	for _, diffs := range d.diffsPerReport {
		firstBad, count := FindAndCount(diffs, func(diff int) bool {
			return !(diff > 0 == (diffs[0] > 0) && diff != 0 && diff < 4 && diff > -4)
		})
		if count == 0 {
			sum++
			continue
		}

		// Either need to remove the element after the diff, or the one before.
		if count <= 2 {
			diffsCopy1 := make([]int, len(diffs))
			diffsCopy2 := make([]int, len(diffs))
			copy(diffsCopy1, diffs)
			copy(diffsCopy2, diffs)
			if firstBad < len(diffs)-1 {
				diffsCopy1[firstBad+1] = diffsCopy1[firstBad+1] + diffsCopy1[firstBad]
			}
			if firstBad > 0 {
				diffsCopy2[firstBad-1] = diffsCopy2[firstBad-1] + diffsCopy2[firstBad]
			}
			condition := func(diff int) bool {
				return diff > 0 == (diffs[0] > 0) && diff != 0 && diff < 4 && diff > -4
			}
			if AllExcept(diffsCopy1, firstBad, condition) || AllExcept(diffsCopy2, firstBad, condition) {
				sum++
				continue
			}
		}

		// Need to remove the first element
		if count >= len(diffs)-2 {
			diffsCopy := make([]int, len(diffs))
			copy(diffsCopy, diffs)
			diffsCopy[1] = diffsCopy[1] + diffsCopy[0]
			diffsCopy = diffsCopy[1:]
			diffs = diffs[1:]

			condition1 := func(diff int) bool {
				return diff > 0 == (diffs[0] > 0) && diff != 0 && diff < 4 && diff > -4
			}

			condition2 := func(diff int) bool {
				return diff > 0 == (diffsCopy[0] > 0) && diff != 0 && diff < 4 && diff > -4
			}

			if All(diffs, condition1) || All(diffsCopy, condition2) {
				sum++
			}
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
