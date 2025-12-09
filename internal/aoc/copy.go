package aoc

func CopyMap[T comparable, Y any](v map[T]Y) map[T]Y {
	result := map[T]Y{}
	for key, value := range v {
		result[key] = value
	}
	return result
}
