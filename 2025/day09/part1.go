package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

func Part1(lines []string) int {
	state := NewState(lines)

	state.Print(nil)

	largestArea := 0

	for _, v1 := range state.Vectors {
		for _, v2 := range state.Vectors {
			if area := Area(v1, v2); area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

type Value string

const (
	Unknown Value = ""
	Blank   Value = "."
	Green   Value = "X"
	Red     Value = "#"
)

type Vector struct {
	X int
	Y int
}

func Area(v1 Vector, v2 Vector) int {
	return int((math.Abs(float64(v2.X)-float64(v1.X)) + 1) * (math.Abs(float64(v2.Y)-float64(v1.Y)) + 1))
}

func ParseVectors(lines []string) []Vector {
	v := []Vector{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		vector := Vector{}
		for i, part := range parts {
			number, err := strconv.Atoi(part)
			aoc.PanicOnErr(err)
			if i == 0 {
				vector.X = number
			} else {
				vector.Y = number
			}
		}

		v = append(v, vector)
	}

	return v
}

type State struct {
	Values  map[Vector]Value
	Vectors []Vector
	Height  int
	Width   int
}

func NewState(lines []string) State {
	state := State{
		Values:  map[Vector]Value{},
		Vectors: ParseVectors(lines),
	}

	// Plot points, and track height and width
	for _, v := range state.Vectors {
		if v.X > state.Width {
			state.Width = v.X + 1
		}

		if v.Y > state.Height {
			state.Height = v.Y + 1
		}

		state.Values[v] = Red
	}

	// Draw lines
	for i, v := range state.Vectors {
		previousIndex := i - 1

		if previousIndex < 0 {
			previousIndex = len(state.Vectors) - 1
		}
		previousV := state.Vectors[previousIndex]

		for _, vector := range drawLine(previousV, v) {
			if vector == previousV || vector == v {
				continue
			}
			state.Values[vector] = Green
		}
	}

	return state
}

func (state State) Print(mod func(state map[Vector]Value) map[Vector]Value) {
	// Don't print large datasets
	if len(state.Vectors) > 100 {
		return
	}

	tmp := aoc.CopyMap(state.Values)
	if mod != nil {
		tmp = mod(tmp)
	}

	for y := 0; y <= state.Height; y++ {
		for x := 0; x <= state.Width; x++ {
			v := tmp[Vector{X: x, Y: y}]
			if v == "" {
				v = "."
			}
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
