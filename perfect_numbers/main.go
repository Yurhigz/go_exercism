package perfect

// Define the Classification type here.
import "errors"

type Classification int

var ErrOnlyPositive = errors.New("input must be a positive integer")

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

func Classify(n int64) (Classification, error) {

	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	v := int(n)
	l := 0
	for i := 1; i < v; i++ {
		if v%i == 0 {
			l += i
		}
	}
	if l > v {
		return ClassificationAbundant, nil
	} else if l == v {
		return ClassificationPerfect, nil
	} else {
		return ClassificationDeficient, nil
	}
}
