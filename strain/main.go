package strain

//"errors"

// appliquer le prédicat à chacun des éléments de la liste et retourner si keep la condition vrai, et discard si la condition est fausse

func Keep[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}

func Discard[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, e := range s {
		if !predicate(e) {
			result = append(result, e)
		}
	}
	return result
}
