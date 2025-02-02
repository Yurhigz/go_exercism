package main

import (
	"fmt"
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	nest := reflect.ValueOf(nested)
	var results []interface{}
	if nest.Kind() == reflect.Array || nest.Kind() == reflect.Slice {
		for j := 0; j < nest.Len(); j++ {
			elem := nest.Index(j).Interface()
			if reflect.ValueOf(elem).Kind() == reflect.Slice {
				results = append(results, Flatten(elem)...)
			} else {
				results = append(results, elem)
			}

		}
	} else {
		results = append(results, nested)
	}
	return results
}

func main() {
	fmt.Println(Flatten([]interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8}))
}
