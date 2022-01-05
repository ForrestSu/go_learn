package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	weekGap = 22
)

var dateRE = regexp.MustCompile(`(\d+)`)

func getYear(week int) int {
	if week < weekGap {
		return 2015
	}
	return 2016
}

// FormatTitle 10月(23-29) 62 => "## 2015-10-23 ~ 2015-10-29 (Week 62)"
func FormatTitle(name string) (string, error) {
	results := dateRE.FindAllString(name, -1)
	if len(results) != 4 {
		return "", fmt.Errorf("%s is mot match, result=%v", name, results)
	}
	list := strsToInts(results)
	var (
		monthStart = list[0]
		dayStart   = list[1]
		dayEnd     = list[2]
		week       = list[3]
	)
	var year = getYear(week)
	var monthEnd = monthStart
	if dayEnd < dayStart {
		monthEnd = monthEnd + 1
	}
	return fmt.Sprintf("%d-%02d-%02d ~ %d-%02d-%02d (Week %d)",
		year, monthStart, dayStart,
		year, monthEnd, dayEnd,
		week), nil
}

func strsToInts(strs []string) []int {
	var ints []int
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil || n <= 0 {
			panic(err)
		}
		ints = append(ints, n)
	}
	return ints
}
