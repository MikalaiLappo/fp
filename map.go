package fp

func Map[T any, R any](input []T, transformer func(arg T) R) []R {
	result := make([]R, len(input))
	for i, x := range input {
		result[i] = transformer(x)
	}
	return result
}
