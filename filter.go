package fp

func Filter[T any](input []T, predicate func(i int, arg T) bool) []T {
	result := make([]T, len(input))
	idx := 0
	for i, x := range input {
		if !predicate(i, x) {
			continue
		}
		result[idx] = x
		idx += 1
	}
	return result[0:idx]
}
