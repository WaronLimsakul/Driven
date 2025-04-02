package tasks

import (
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
)

var weekDaysArray [7]string = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

// Return the monday (start) and sunday (end) of the week today is in
func GetWeekRange(today time.Time) (monday, sunday time.Time) {
	goBack := GetWeekDayNum(today)

	monday = today.AddDate(0, 0, -goBack)
	sunday = today.AddDate(0, 0, 6-goBack)
	return monday, sunday
}

// Get tasks in a week and return [][]task. Start by monday.
// Not guarantee that all inner slices are not nil.
// time: O(n)
func GroupTaskDate(tasks []database.Task) [][]database.Task {
	week := make([][]database.Task, 7)
	for _, task := range tasks {
		i := GetWeekDayNum(task.Date)
		week[i] = append(week[i], task)
	}
	return week
}

// Monday is 0, until Sunday is 6
func GetWeekDayNum(day time.Time) int {
	i := int(day.Weekday()) - 1 //sunday is 0
	if i < 0 {
		i = 6
	}
	return i
}

// Shorten week day name to just 3 characters
func GetWeekDayStr(day time.Time) string {
	dayNum := GetWeekDayNum(day)
	return weekDaysArray[dayNum]
}

// get monday (any time) and return the string
// slice for date in dd/mm format
func SpanWeekDateByFormat(monday time.Time, format func(time.Time) string) []string {
	week := []string{}
	cur := monday
	for range 7 {
		week = append(week, format(cur))
		cur = cur.Add(24 * time.Hour)
	}
	return week
}

// return a slice with 7 days start with monday (it can be anyday though)
func SpanWeekDate(monday time.Time) []time.Time {
	week := []time.Time{}
	cur := monday
	for range 7 {
		week = append(week, cur)
		cur = cur.Add(24 * time.Hour)
	}
	return week
}
