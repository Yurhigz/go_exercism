package lsproduct
import (
    "errors"

)
func LargestSeriesProduct(digits string, span int) (int64, error) {
    value := int64(0)

    if span <= 0 {
        return 0, errors.New("span must be greater than 0")
    }
    if span > len(digits) {
        return 0, errors.New("span must be smaller than string length")
    }
	for _,char := range digits {
        if char < '0' || char > '9' {
            return 0, errors.New("digits input must only contain digits")
        }
    }
    for i:=0 ; i<= len(digits)-span;i++ {
        product := int64(1) 
        for j:=0 ; j < span ; j++ {
            digit := digits[i+j] - '0'
            product *= int64(digit)
        }
        if value < product {
            value = product
        }
    }
	return value,nil
}
