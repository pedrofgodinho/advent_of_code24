package main

import (
	"flag"
	"fmt"
	"github.com/pedrofgodinho/advent_of_code24/solver"
	"io"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	dayFlag := flag.Int("day", 0, "Day to run")
	inputDir := flag.String("inputs", "inputs", "Input directory")
	flag.Parse()

	// results array
	var results []DayStats

	switch {
	case *dayFlag == 0:
		// Run all days
		results_chan := make(chan DayStats, 25)
		for i := 1; i <= len(solver.Solvers()); i++ {
			go func(i int) {
				result, err := solveDay(*inputDir, i)
				if err != nil {
					fmt.Printf("Error running day %d: %v", i, err)
				}
				results_chan <- result
			}(i)
		}
		for i := 1; i <= len(solver.Solvers()); i++ {
			results = append(results, <-results_chan)
		}
		// Close the channel
		close(results_chan)
		sort.Sort(byDay(results))
	case *dayFlag < 0 || *dayFlag > len(solver.Solvers()):
		// Invalid day
		fmt.Printf("Invalid day: %d", *dayFlag)
	default:
		result, err := solveDay(*inputDir, *dayFlag)
		if err != nil {
			fmt.Printf("Error running day %d: %v", *dayFlag, err)
		}
		results = append(results, result)
	}

	fmt.Println(asTable(results))
}

type DayStats struct {
	day         int
	part1Time   time.Duration
	part2Time   time.Duration
	part1Result string
	part2Result string
	parsed      bool
	parseTime   time.Duration
}

type byDay []DayStats

func (a byDay) Len() int           { return len(a) }
func (a byDay) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDay) Less(i, j int) bool { return a[i].day < a[j].day }

// Get the input for a given day
func getInput(dir string, day int) (string, error) {
	filePath := fmt.Sprintf("%s/day%d.txt", dir, day)
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

// Solve a given day
func solveDay(inputDir string, day int) (DayStats, error) {
	// Get the input
	input, err := getInput(inputDir, day)
	if err != nil {
		return DayStats{}, err
	}

	s := solver.Solvers()[day-1]

	stats := DayStats{day: day}

	start := time.Now()
	stats.parsed, err = s.Parse(input)
	if err != nil {
		return DayStats{}, err
	}
	stats.parseTime = time.Since(start)

	start = time.Now()
	stats.part1Result, err = s.Part1()
	if err != nil {
		return DayStats{}, err
	}
	stats.part1Time = time.Since(start)

	start = time.Now()
	stats.part2Result, err = s.Part2()
	if err != nil {
		return DayStats{}, err
	}
	stats.part2Time = time.Since(start)

	return stats, nil
}

func asTable(results []DayStats) string {
	// Day | Parse Time | Part 1 Time | Part 2 Time | Part 1 Result | Part 2 Result
	colWidths := make([]int, 6)
	days := make([]string, len(results)+1)
	parseTimes := make([]string, len(results)+1)
	part1Times := make([]string, len(results)+1)
	part2Times := make([]string, len(results)+1)
	part1Results := make([]string, len(results)+1)
	part2Results := make([]string, len(results)+1)

	days[0] = "Day"
	parseTimes[0] = "Parse Time"
	part1Times[0] = "Part 1 Time"
	part2Times[0] = "Part 2 Time"
	part1Results[0] = "Part 1 Result"
	part2Results[0] = "Part 2 Result"

	for i, result := range results {
		days[i+1] = fmt.Sprintf("%d", result.day)
		if !result.parsed {
			parseTimes[i+1] = "N/A"
		} else {
			parseTimes[i+1] = fmt.Sprintf("%v", result.parseTime)
		}
		part1Times[i+1] = fmt.Sprintf("%v", result.part1Time)
		part2Times[i+1] = fmt.Sprintf("%v", result.part2Time)
		part1Results[i+1] = result.part1Result
		part2Results[i+1] = result.part2Result
	}

	// Calculate column widths
	compareLengths := func(a, b string) int {
		len1 := utf8.RuneCountInString(a)
		len2 := utf8.RuneCountInString(b)
		return len1 - len2
	}
	colWidths[0] = len(slices.MaxFunc(days, compareLengths))
	colWidths[1] = len(slices.MaxFunc(parseTimes, compareLengths))
	colWidths[2] = len(slices.MaxFunc(part1Times, compareLengths))
	colWidths[3] = len(slices.MaxFunc(part2Times, compareLengths))
	colWidths[4] = len(slices.MaxFunc(part1Results, compareLengths))
	colWidths[5] = len(slices.MaxFunc(part2Results, compareLengths))

	// Print header
	var output strings.Builder
	output.WriteString("┌─")
	lines := []string{
		strings.Repeat("─", colWidths[0]),
		strings.Repeat("─", colWidths[1]),
		strings.Repeat("─", colWidths[2]),
		strings.Repeat("─", colWidths[3]),
		strings.Repeat("─", colWidths[4]),
		strings.Repeat("─", colWidths[5]),
	}
	output.WriteString(strings.Join(lines, "─┬─"))
	output.WriteString("─┐\n")
	for i := 0; i < len(days); i++ {
		output.WriteString("│ ")
		output.WriteString(strings.Repeat(" ", colWidths[0]-utf8.RuneCountInString(days[i])))
		output.WriteString(days[i])
		output.WriteString(" │ ")
		output.WriteString(strings.Repeat(" ", colWidths[1]-utf8.RuneCountInString(parseTimes[i])))
		output.WriteString(parseTimes[i])
		output.WriteString(" │ ")
		output.WriteString(strings.Repeat(" ", colWidths[2]-utf8.RuneCountInString(part1Times[i])))
		output.WriteString(part1Times[i])
		output.WriteString(" │ ")
		output.WriteString(strings.Repeat(" ", colWidths[3]-utf8.RuneCountInString(part2Times[i])))
		output.WriteString(part2Times[i])
		output.WriteString(" │ ")
		output.WriteString(strings.Repeat(" ", colWidths[4]-utf8.RuneCountInString(part1Results[i])))
		output.WriteString(part1Results[i])
		output.WriteString(" │ ")
		output.WriteString(strings.Repeat(" ", colWidths[5]-utf8.RuneCountInString(part2Results[i])))
		output.WriteString(part2Results[i])
		output.WriteString(" │\n")
	}
	output.WriteString("└─")
	output.WriteString(strings.Join(lines, "─┴─"))
	output.WriteString("─┘\n")

	return output.String()
}
