package day09

import (
	"slices"
)

func Part2(lines []string) int {
	state := NewState(lines)
	yMap := map[int][]int{}

	// Draw lines
	for i, v := range state.Vectors {
		previousIndex := i - 1
		yMap[v.X] = append(yMap[v.X], v.Y)

		if previousIndex < 0 {
			previousIndex = len(state.Vectors) - 1
		}
		previousV := state.Vectors[previousIndex]
		if v.X == previousV.X {
			y1 := v.Y
			y2 := previousV.Y
			if y1 > y2 {
				y2tmp := y1
				y1 = y2
				y2 = y2tmp
			}
			for i := y1 + 1; i < y2; i++ {
				state.Values[Vector{X: v.X, Y: i}] = Green
				yMap[v.X] = append(yMap[v.X], i)
			}
		} else if v.Y == previousV.Y {
			x1 := v.X
			x2 := previousV.X
			if x1 > x2 {
				x2tmp := x1
				x1 = x2
				x2 = x2tmp
			}
			for i := x1 + 1; i < x2; i++ {
				state.Values[Vector{X: i, Y: v.Y}] = Green
				yMap[i] = append(yMap[i], v.Y)
			}
		}
	}

	// Sort the map
	for x := range yMap {
		slices.Sort(yMap[x])
	}

	// Print starting state
	state.Print(nil)

	isVectorInPolygon := func(vector Vector) bool {
		existing := state.Values[vector]
		if existing != "" {
			return existing == Green || existing == Red
		}
		result := valuesLessThan(vector, yMap[vector.X])%2 != 0
		if result {
			state.Values[vector] = Green
		} else {
			state.Values[vector] = Blank
		}
		return result
	}

	isLineInPolygon := func(start Vector, end Vector) bool {
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

			if !isVectorInPolygon(vector) {
				return false
			}

			if vector == end {
				break
			}
		}

		return true
	}

	isPolygonInPolygon := func(polygon []Vector) bool {
		for i, vector1 := range polygon {
			vector2Index := i + 1
			if vector2Index > len(polygon)-1 {
				vector2Index = 0
			}

			vector2 := polygon[vector2Index]

			if !isLineInPolygon(vector1, vector2) {
				return false
			}
		}

		return true
	}

	largestArea := 0
	for _, v1 := range state.Vectors {
		for _, v2 := range state.Vectors {
			if v1 == v2 {
				continue
			}

			// targetX := Vector{X: 9, Y: 5}
			// targetY := Vector{X: 2, Y: 3}
			// if v1 != targetX || v2 != targetY {
			// 	continue
			// }

			v3 := Vector{X: v1.X, Y: v2.Y}
			v4 := Vector{X: v2.X, Y: v1.Y}

			polygon := []Vector{
				v1, v3, v2, v4,
			}

			state.Print(func(state map[Vector]Value) map[Vector]Value {
				state[v1] = "A"
				state[v2] = "A"
				state[v3] = "B"
				state[v4] = "B"
				return state
			})

			if !isPolygonInPolygon(polygon) {
				continue
			}

			area := Area(v1, v2)
			if area > largestArea {
				state.Print(func(state map[Vector]Value) map[Vector]Value {
					state[v1] = "A"
					state[v2] = "A"
					state[v3] = "B"
					state[v4] = "B"
					return state
				})
				largestArea = area
			}
		}
	}

	return largestArea
}

func valuesLessThan(vector Vector, values []int) int {
	value := vector.Y
	count := 0
	for _, v := range values {
		if v >= value {
			count++
		}
	}

	return count
}
