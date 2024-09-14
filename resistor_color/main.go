package resistorcolor

import "strings"

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, v := range Colors() {
		if strings.ToLower(color) == v {
			return i
		}
	}
	return -1
}
