package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type Day5 struct {
	rules map[int][]int // Before -> After
	pages [][]int
}

func (d *Day5) Parse(input string) (bool, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(parts) != 2 {
		return false, fmt.Errorf("expected 2 blocks, got %d", len(parts))
	}
	rules := strings.Split(parts[0], "\n")
	pages := strings.Split(parts[1], "\n")

	d.rules = make(map[int][]int)

	for _, rule := range rules {
		before := 0
		after := 0
		_, err := fmt.Sscanf(rule, "%d|%d", &before, &after)
		if err != nil {
			return false, err
		}

		d.rules[before] = append(d.rules[before], after)
	}

	for _, page := range pages {
		var p []int
		for _, num := range strings.Split(page, ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				return false, err
			}
			p = append(p, n)
		}
		d.pages = append(d.pages, p)
	}
	return true, nil
}

func (d *Day5) Part1() (string, error) {
	sum := 0
PagesLoop:
	for _, page := range d.pages {
		seenPages := make(map[int]bool)
		for _, num := range page {
			if pagesAfter, ok := d.rules[num]; ok {
				// There is a rule for this page before another
				for _, after := range pagesAfter {
					if _, ok := seenPages[after]; ok {
						// The page that should go after is before
						continue PagesLoop
					}
				}
			}

			seenPages[num] = true
		}

		sum += page[len(page)/2]
	}
	return fmt.Sprintf("%d", sum), nil
}

func (d *Day5) Part2() (string, error) {
	// This seems like a particularly messy way to solve this problem
	// Not to mention slow
	// But it works

	sum := 0
	for _, page := range d.pages {
		fixed := false
		pageGood := false
	PageGood:
		for !pageGood {
			pageGood = true
			seenPages := make(map[int]int)
			for i, num := range page {
				if pagesAfter, ok := d.rules[num]; ok {
					// There is a rule for this page before another
					for _, after := range pagesAfter {
						if badIndex, ok := seenPages[after]; ok {
							// The page that should go after is before
							// Need to remove that page and add it after
							for j := badIndex; j < i; j++ {
								page[j] = page[j+1]
							}
							page[i] = after
							fixed = true
							pageGood = false
							continue PageGood
						}
					}
				}

				seenPages[num] = i
			}
		}

		if fixed {
			sum += page[len(page)/2]
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
