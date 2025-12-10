package day09

import (
	"log"
	"time"
)

func Part2(lines []string) int {
	state := NewState(lines)
	// Print starting state
	state.Print(nil)

	largestArea := 0
	totalOptions := len(state.Vectors) * len(state.Vectors)
	_ = totalOptions

	timer := &Timer{
		Current:   0,
		Total:     len(state.Vectors) * len(state.Vectors),
		StartTime: time.Now(),
	}
	_ = timer
	for _, v1 := range state.Vectors {
		for _, v2 := range state.Vectors {
			// timer.Tick()

			// Skip checking self
			if v1 == v2 {
				continue
			}

			// Calculated other two points to make rectangle
			v3 := Vector{X: v1.X, Y: v2.Y}
			v4 := Vector{X: v2.X, Y: v1.Y}

			// Early skip if extrapolated points not in polygon
			if !state.IsVectorInPolygon(v3) {
				continue
			}
			if !state.IsVectorInPolygon(v4) {
				continue
			}

			polygon := []Vector{
				v1, v3, v2, v4,
			}

			if !state.isPolygonInPolygon(polygon) {
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

func (state State) IsVectorInPolygon(vector Vector) bool {
	// check for cache
	existing := state.Values[vector]
	if existing != "" {
		return existing == Green || existing == Red
	}

	// Calculate
	result := state.pointInPolygon(vector)

	// Save result to cache
	if result {
		state.Values[vector] = Green
	} else {
		state.Values[vector] = "-"
	}

	return result
}

func (state State) pointInPolygon(pt Vector) bool {
	inside := false
	n := len(state.Vectors)

	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := state.Vectors[i].X, state.Vectors[i].Y
		xj, yj := state.Vectors[j].X, state.Vectors[j].Y

		intersect := ((yi > pt.Y) != (yj > pt.Y)) &&
			(pt.X < (xj-xi)*(pt.Y-yi)/(yj-yi)+xi)
		if intersect {
			inside = !inside
		}
	}
	return inside
}

func (state State) isLineInPolygon(start Vector, end Vector) bool {
	for _, vector := range drawLine(start, end) {
		if !state.IsVectorInPolygon(vector) {
			return false
		}
	}

	return true
}

func (state State) isPolygonInPolygon(polygon []Vector) bool {
	for i, vector1 := range polygon {
		vector2Index := i + 1
		if vector2Index > len(polygon)-1 {
			vector2Index = 0
		}

		vector2 := polygon[vector2Index]
		if !state.isLineInPolygon(vector1, vector2) {
			return false
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

type Timer struct {
	Current   int
	Total     int
	StartTime time.Time
}

func (t *Timer) Tick() {
	t.Current++
	elapsedTime := time.Since(t.StartTime)
	perLoopAvg := elapsedTime / time.Duration(t.Current)
	log.Printf(
		"%d/%d time remaining: %s - per second: %d",
		t.Current,
		t.Total,
		perLoopAvg*time.Duration(t.Total-t.Current),
		time.Second/perLoopAvg,
	)
}
