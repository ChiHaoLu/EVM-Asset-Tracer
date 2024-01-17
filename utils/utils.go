package utils

func Filter[T any](elems []T, filter func(elem T) bool) []T {
	var res []T
	for _, elem := range elems {
		if filter(elem) {
			res = append(res, elem)
		}
	}
	return res
}