package _1154_DayOfYear

import "time"

func dayofYear(target string) int {

	parseTime, _ := time.Parse("2006-01-02", target)
	isLeapYear := isLeapYear(parseTime.Year())
	monthsDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	ans := 0
	for i := 0; i < int(parseTime.Month())-1; i++ {
		ans += monthsDays[i]
	}
	if isLeapYear {
		ans += 1
	}
	return ans + parseTime.Day()
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
