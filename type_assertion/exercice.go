package sorting
import ("fmt"
		"strconv"
        )
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {

	return fmt.Sprintf("This is the number %.1f",f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f",float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	str, ok := fnb.(FancyNumber)
    if ok {
        value,_ := strconv.Atoi(str.Value())
        return value
    }
    return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	if ExtractFancyNumber(fnb) == 0 {
        return "This is a fancy box containing the number 0.0"
    } else {
        return fmt.Sprintf("This is a fancy box containing the number %.1f",float64(ExtractFancyNumber(fnb)))
    }
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
        case int:
			return DescribeNumber(float64(v))
        case float64:
        	return DescribeNumber(v)
        case NumberBox : 
			return DescribeNumberBox(v)
        case FancyNumberBox:
        	return DescribeFancyNumberBox(v)
        default :
        	return "Return to sender"
    }
} 
