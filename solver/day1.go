package solver

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct {
	list1            []int
	list2            []int
	list2Occurrences map[int]int
}

func (d *Day1) Parse(input string) (bool, error) {
	d.list2Occurrences = make(map[int]int)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return false, fmt.Errorf("invalid input: %s", line)
		}
		first, err := strconv.Atoi(fields[0])
		if err != nil {
			return false, err
		}
		second, err := strconv.Atoi(fields[1])
		if err != nil {
			return false, err
		}
		d.list1 = append(d.list1, first)
		d.list2 = append(d.list2, second)
		d.list2Occurrences[second]++
	}
	slices.Sort(d.list1)
	slices.Sort(d.list2)
	return true, nil
}

func (d *Day1) Part1() (string, error) {
	sum := 0
	for i := 0; i < len(d.list1); i++ {
		diff := d.list2[i] - d.list1[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return fmt.Sprintf("%d", sum), nil
}

func (d *Day1) Part2() (string, error) {
	sum := 0
	for i := 0; i < len(d.list1); i++ {
		score := d.list1[i] * d.list2Occurrences[d.list1[i]]
		sum += score
	}
	return fmt.Sprintf("%d", sum), nil
}
