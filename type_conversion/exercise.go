package expenses
import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func predicate(r Record) bool {
    return r.Day == 1
}
func Filter(in []Record, predicate func(Record) bool) []Record {
	records := []Record{}
    for _,v := range in {
        if predicate(v){ 
        	records = append(records,v)
    	}
    }
	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
        if r.Day <= p.To && r.Day >= p.From {
            return true
        }
        return false
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
    return func(r Record) bool {
        if r.Category == c {
            return true
        }
        return false
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    filtered := Filter(in,ByDaysPeriod(p))
    var totalAmount float64 = 0 
    for _,v := range filtered {
        totalAmount = totalAmount + v.Amount
    }
	return totalAmount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
    err := errors.New("unknown category")
	CategoryFiltered := Filter(in,ByCategory(c))
    if len(CategoryFiltered) == 0 {
        return 0,err
    }
	return TotalByPeriod(CategoryFiltered,p), nil
}
