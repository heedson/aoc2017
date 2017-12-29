package day2

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/heedson/aoc2017/common"
)

// Day is the implementation of day 2.
type Day struct {
	rows []row
}

type row struct {
	values []int
}

func (r row) diff() int {
	lowest := math.MaxInt64
	highest := math.MinInt64
	for _, v := range r.values {
		if v > highest {
			highest = v
		}

		if v < lowest {
			lowest = v
		}
	}
	return highest - lowest
}

func (r row) div() int {
	for i, v := range r.values {
		for j, d := range r.values {
			if i == j {
				continue
			}

			if v%d == 0 {
				return v / d
			}
		}
	}
	panic("no clean divides")
}

// New returns a new instantiation of Day.
func New() Day {
	reader, err := os.Open("./day2/day2.txt")
	if err != nil {
		panic(err)
	}

	rows, err := getRows(reader)
	if err != nil {
		panic(err)
	}

	return Day{
		rows: rows,
	}
}

func getRows(reader io.Reader) ([]row, error) {
	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(input, "\n")

	rows := make([]row, 0, len(lines))

	for _, l := range lines {
		var row row

		values := strings.Split(l, "\t")

		row.values = make([]int, 0, len(values))

		for _, v := range values {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			row.values = append(row.values, n)
		}

		rows = append(rows, row)
	}

	return rows, nil
}

// RunP1 runs the task's Part 1 for the day. It prints the challenge output.
func (d Day) RunP1() error {
	var total int
	for _, r := range d.rows {
		total += r.diff()
	}

	fmt.Println(total)

	return nil
}

// RunP2 runs the task's Part 2 for the day. It prints the challenge output.
func (d Day) RunP2() error {
	var total int
	for _, r := range d.rows {
		total += r.div()
	}

	fmt.Println(total)

	return nil
}
