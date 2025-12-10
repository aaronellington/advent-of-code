package day09

import (
	"log"
	"sort"
)

var expensiveCalculations = 0

func Part2(lines []string) int {
	state := NewState(lines)
	state.Print()

	options := []Option{}
	for _, v1 := range state.Vectors {
		for _, v2 := range state.Vectors {
			if v1 == v2 {
				continue
			}
			options = append(options, Option{
				Start: v1,
				End:   v2,
				Area:  Area(v1, v2),
			})
		}
	}
	sort.Slice(options, func(i, j int) bool {
		return options[i].Area > options[j].Area
	})

	log.Println("starting")
	largestArea := 0
	for _, option := range options {
		v1 := option.Start
		v2 := option.End

		// Calculated other two points to make rectangle
		v3 := Vector{X: v1.X, Y: v2.Y}
		v4 := Vector{X: v2.X, Y: v1.Y}

		// Early skip if extrapolated points not in polygon
		if !state.IsVectorInPolygon(v4) || !state.IsVectorInPolygon(v3) {
			continue
		}

		// Skip because not valid
		if !state.isPolygonInPolygon([]Vector{
			v1, v3, v2, v4,
		}) {
			continue
		}

		state.Values[v1] = "1"
		state.Values[v2] = "2"
		state.Values[v3] = "3"
		state.Values[v4] = "4"
		state.Print()

		largestArea = option.Area
		log.Printf("Found answer after %d expensive calculations", expensiveCalculations)
		break
	}

	return largestArea
}

func (state State) IsVectorInPolygon(vector Vector) bool {
	// check for cache
	existing := state.Values[vector]
	if existing != "" {
		return existing == Green || existing == Red
	}

	// Calculate
	result := state.pointInPolygon(vector)
	expensiveCalculations++

	// Save result to cache
	if result {
		state.Values[vector] = Green
	} else {
		state.Values[vector] = "-"
	}

	return result
}

func (state State) pointInPolygon(vector Vector) bool {
	inside := false
	n := len(state.Vectors)

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := state.Vectors[i].X, state.Vectors[i].Y
		xj, yj := state.Vectors[j].X, state.Vectors[j].Y

		intersect := ((yi > vector.Y) != (yj > vector.Y)) &&
			(vector.X < (xj-xi)*(vector.Y-yi)/(yj-yi)+xi)
		if intersect {
			inside = !inside
		}
	}
	return inside
}

func (state State) isPolygonInPolygon(polygon []Vector) bool {
	for i, start := range polygon {
		vector2Index := i + 1
		if vector2Index > len(polygon)-1 {
			vector2Index = 0
		}

		end := polygon[vector2Index]
		for _, vector := range drawLine(start, end) {
			if !state.IsVectorInPolygon(vector) {
				return false
			}
		}
	}

	return true
}

func drawLine(start Vector, end Vector) []Vector {
	vectors := []Vector{}
	xMove := 0
	if start.X < end.X {
		xMove = 1
	} else if start.X > end.X {
		xMove = -1
	}
	yMove := 0
	if start.Y < end.Y {
		yMove = 1
	} else if start.Y > end.Y {
		yMove = -1
	}

	i := 0
	for {
		vector := Vector{
			X: start.X + xMove*i,
			Y: start.Y + yMove*i,
		}
		i++

		vectors = append(vectors, vector)

		if vector == end {
			break
		}
	}

	return vectors
}

type Option struct {
	Start Vector
	End   Vector
	Area  int
}
