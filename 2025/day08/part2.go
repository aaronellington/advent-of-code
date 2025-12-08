package day08

import "sort"

type Part2 struct {
	MaxConnections int
}

func (part2 Part2) Solve(lines []string) int {
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
		}

		sort.Slice(d, func(i, j int) bool {
			return d[i].Distance < d[j].Distance
		})

		return d
	}()

	lastDistance := distances[0]
	for i, d := range distances {
		if len(vectors) == len(assignments) {
			lastDistance = distances[i-1]
			break
		}

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
			// circuits[existingV2Circuit] = []VectorID{} // Delete v2Circuit
			delete(circuits, existingV2Circuit)
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

	return int(vectors[lastDistance.V1].X * vectors[lastDistance.V2].X)
}
