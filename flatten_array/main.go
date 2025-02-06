package main

import (
	"fmt"
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	nest := reflect.ValueOf(nested)
	results := make([]interface{}, 0)

	if nest.Len() == 0 && nest.Kind() == reflect.Slice {
		return results
	}

	if nest.Kind() == reflect.Array || nest.Kind() == reflect.Slice {

		for j := 0; j < nest.Len(); j++ {

			elem := nest.Index(j).Interface()

			elemValue := reflect.ValueOf(elem)

			if elem == nil {
				continue
			}
			if elemValue.Kind() == reflect.Slice {
				if elemValue.Len() == 0 {
					continue
				}

				flattened := Flatten(elem)
				if len(flattened) > 0 {
					results = append(results, flattened...)
				}
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
	// fmt.Println(Flatten([]interface{}{1, 2, interface{}(nil)}))
	// fmt.Println(Flatten([]interface{}{[]interface{}{[]interface{}{}}}) == nil)
	// fmt.Println(Flatten([]interface{}{[]interface{}{[]interface{}{}}}))
}
