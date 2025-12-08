package day08

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/aaronellington/advent-of-code/internal/aoc"
)

type Distance struct {
	V1       VectorID
	V2       VectorID
	Distance float64
}

type Part1 struct {
	MaxConnections int
}

func (part1 Part1) Solve(lines []string) int {
	vectors := toVectors(lines)
	assignments := map[VectorID]CircuitID{}
	circuits := map[CircuitID][]VectorID{}

	// Get all distances
	distances := func() []Distance {
		d := []Distance{}
		for i1, v1 := range vectors {
			for i2, v2 := range vectors {
				if i1 >= i2 {
					continue
				}
				d = append(d, Distance{
					V1:       i1,
					V2:       i2,
					Distance: distance(v1, v2),
				})
			}
			sort.Slice(d, func(i, j int) bool {
				return d[i].Distance < d[j].Distance
			})
			if len(d) > part1.MaxConnections {
				d = d[0:part1.MaxConnections]
			}
		}

		return d
	}()

	for i := 0; i < part1.MaxConnections; i++ {
		d := distances[i]

		existingV1Circuit := assignments[d.V1]
		existingV2Circuit := assignments[d.V2]

		// Both are already part of a circuit
		if existingV1Circuit > 0 && existingV2Circuit > 0 {
			// On same circuit, nothing happens
			if existingV1Circuit == existingV2Circuit {
				continue
			}

			// On different circuits, Merge circuits
			for _, v2ID := range circuits[existingV2Circuit] {
				assignments[v2ID] = existingV1Circuit
				circuits[existingV1Circuit] = append(circuits[existingV1Circuit], v2ID)
			}
			circuits[existingV2Circuit] = []VectorID{} // Delete v2Circuit
			continue
		}

		if existingV1Circuit > 0 {
			// v2 not in a circuit
			assignments[d.V2] = existingV1Circuit
			circuits[existingV1Circuit] = append(circuits[existingV1Circuit], d.V2)
			continue
		} else if existingV2Circuit > 0 {
			// v1 not in a circuit
			assignments[d.V1] = existingV2Circuit
			circuits[existingV2Circuit] = append(circuits[existingV2Circuit], d.V1)
			continue
		}

		// Create new circuit
		newCircuitID := CircuitID(len(circuits) + 1)
		circuits[newCircuitID] = []VectorID{d.V1, d.V2}
		assignments[d.V1] = newCircuitID
		assignments[d.V2] = newCircuitID
	}

	// Sort the circuits by size
	circuitArray := func() []Circuit {
		x := []Circuit{}

		for circuitID, vectorIDs := range circuits {
			x = append(x, Circuit{
				CircuitID: circuitID,
				VectorIDs: vectorIDs,
			})
		}
		sort.Slice(x, func(c1 int, c2 int) bool {
			return len(x[c1].VectorIDs) > len(x[c2].VectorIDs)
		})

		return x
	}()

	return len(circuitArray[0].VectorIDs) * len(circuitArray[1].VectorIDs) * len(circuitArray[2].VectorIDs)
}

type CircuitID int
type Circuit struct {
	CircuitID CircuitID
	VectorIDs []VectorID
}

type VectorID int
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func distance(p1, p2 Vector3) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func toVectors(lines []string) map[VectorID]Vector3 {
	result := map[VectorID]Vector3{}
	for i, line := range lines {
		v := Vector3{}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		for i, numberString := range parts {
			number, err := strconv.ParseFloat(numberString, 64)
			aoc.PanicOnErr(err)
			switch i {
			case 0:
				v.X = number
			case 1:
				v.Y = number
			case 2:
				v.Z = number
			}
		}

		result[VectorID(i+1)] = v
	}

	return result
}
