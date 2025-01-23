package main

import (
	"fmt"
	"time"
)

// WeekSchedule is an alias for defining dates based on weeks.
type WeekSchedule int

// Aliases for WeekSchedule coerced into actual days of the month.
const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

// Day parses WeekSchedule expressions into actual dates.
func Day(w WeekSchedule, wd time.Weekday, m time.Month, y int) int {
	if w == Last {
		m++
	}
	t := time.Date(y, m, int(w), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(wd-t.Weekday()+7)%7
}

func main() {
	fmt.Println(Day(Teenth, time.Monday, 8, 2024))
	// fmt.Println(Day(Teenth))

}
