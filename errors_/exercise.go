package thefarm

import (
    "errors"
    "fmt"
)
type CowError struct {
	 message string
    cowNumber int
}

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, cowNumber int) (float64, error) {
    fa, err := f.FodderAmount(cowNumber) 
    if err != nil {
        return 0,err
    }

    ff, err := f.FatteningFactor()
    if err != nil {
        return 0,err
    }
    r := fa*ff/float64(cowNumber)
    return r,err
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, cowNumber int) (float64, error) {
    if cowNumber > 0 {
        return DivideFood(f,cowNumber)
    } else {
        var Err = errors.New("invalid number of cows")
        return 0, Err
    }
}
// TODO: define the 'ValidateNumberOfCows' function


func (c *CowError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s",c.cowNumber,c.message)
}

func ValidateNumberOfCows(cowNumber int) error{
    if cowNumber < 0 {
        return &CowError{"there are no negative cows",cowNumber}
    }
    if cowNumber == 0 {
        return &CowError{"no cows don't need food",cowNumber}
    }
    return nil
}


