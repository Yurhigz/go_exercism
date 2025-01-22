package main

import (
	"fmt"
	"time"
)

// Define the WeekSchedule type here.

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, int(wSched), 0, 0, 0, 0, time.UTC).Day()
	return d
}

func main() {
	fmt.Println(Day(Teenth, time.Monday, 8, 2013))
}
