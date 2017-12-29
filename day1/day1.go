package day1

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/heedson/aoc2017/common"
)

// Day is the implementation of day 1.
type Day struct {
	digits []int
}

// New returns a new instantiation of Day.
func New() Day {
	reader, err := os.Open("./day1/day1.txt")
	if err != nil {
		panic(err)
	}

	digits, err := getDigits(reader)
	if err != nil {
		panic(err)
	}

	return Day{
		digits: digits,
	}
}

func getDigits(reader io.Reader) ([]int, error) {
	input, err := common.LoadStringFromReader(reader)
	if err != nil {
		return nil, err
	}

	digits := make([]int, 0, len(input))

	for _, n := range input {
		d, err := strconv.Atoi(string(n))
		if err != nil {
			return nil, err
		}

		digits = append(digits, d)
	}

	return digits, nil
}

// RunP1 runs the task's Part 1 for the day. It prints the challenge output.
func (d Day) RunP1() error {
	var next int
	var total int

	for _, n := range d.digits {
		next++
		if next == len(d.digits) {
			next = 0
		}

		if n == d.digits[next] {
			total += n
		}
	}

	fmt.Println(total)

	return nil
}

// RunP2 runs the task's Part 2 for the day. It prints the challenge output.
func (d Day) RunP2() error {
	lenDigits := len(d.digits)
	if lenDigits%2 != 0 {
		return errors.New("length of input not even")
	}

	var total int

	halfLen := lenDigits / 2
	bottomHalf := d.digits[:halfLen]
	topHalf := d.digits[halfLen:]
	for i, n := range bottomHalf {
		if n == topHalf[i] {
			total += n
		}
	}

	total *= 2

	fmt.Println(total)

	return nil
}
