package main

import (
	"fmt"
	"reflect"
)
func Flatten(nested interface{}) []interface{} {
	nest := reflect.ValueOf(nested)
	if nest.Len() == 0 && nest.Kind() == reflect.Slice {
		return []interface{}{}
	}

	var results []interface{}

	if nest.Kind() == reflect.Array || nest.Kind() == reflect.Slice {
		for j := 0; j < nest.Len(); j++ {
			elem := nest.Index(j).Interface()
			elemValue := reflect.ValueOf(elem)
			if elemValue.Kind() == reflect.Slice && elemValue.Len() > 0 {
				results = append(results, Flatten(elem)...)
			} else if elem != nil {
				results = append(results, elem)
			}
		}
	} else if reflect.ValueOf(nested).Len() > 0 {
		results = append(results, nested)
	} 
    if results == nil {
        return []interface{}{}
    }
	return results
}

func main() {
	fmt.Println(Flatten([]interface{}{1, []interface{}{2, 3, 4, 5, 6, 7}, 8}))
	fmt.Println(Flatten([]interface{}{1, 2, interface{}(nil)}))
	fmt.Println(Flatten([]interface{}{[]interface{}{[]interface{}{}}}) == nil)
}
