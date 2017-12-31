package day3

import (
	"fmt"
	"math"
)

var target = 289326

// Day is the implementation of day 3.
type Day struct {
	target int
}

// New returns a new instantiation of Day.
func New() Day {
	return Day{
		target: target,
	}
}

// RunP1 runs the task's Part 1 for the day. It prints the challenge output.
func (d Day) RunP1() error {
	targetSqrt := math.Sqrt(float64(d.target))

	prevOddSqrt := int(math.Floor(targetSqrt))
	if prevOddSqrt%2 == 0 {
		prevOddSqrt--
	}

	prevCorner := prevOddSqrt * prevOddSqrt
	nextCorner := (prevOddSqrt + 2) * (prevOddSqrt + 2)

	/*
		4    3    2    1    0
		17   16   15   14   13
		18   5    4    3    12
		19   6    1    2    11
		20   7    8    9    10
		21   22   23   24   25  ...
	*/

	// Logic
	sideLength := (nextCorner - prevCorner) / 4

	position := (d.target - prevCorner) % sideLength

	offset := int(math.Abs(float64(position - (sideLength / 2))))

	dist := int(math.Sqrt(float64(nextCorner))-1) / 2

	fmt.Println(dist + offset)

	return nil
}

// RunP2 runs the task's Part 2 for the day. It prints the challenge output.
func (d Day) RunP2() error {
	return nil
}
