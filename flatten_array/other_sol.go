package flatten_array

// Flatten expects an (nested) slice / array and returns the flattened slice
func Flatten(arr interface{}) []interface{} {
	flat := make([]interface{}, 0)

	switch elem := arr.(type) {
	case []interface{}:
		for _, v := range elem {
			flat = append(flat, Flatten(v)...)
		}
	case interface{}:
		flat = append(flat, elem)
	}
	return flat
}
