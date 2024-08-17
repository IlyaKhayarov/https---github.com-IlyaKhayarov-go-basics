package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Friday(1703, 0))
}

func Friday(start, end int) string {
	var ret string
	if end < start {
		end = start
	}
	for y := start; y <= end; y++ {
		for m := 1; m <= 12; m++ {
			lastDayOfMonth := time.Date(y, time.Month(m+1), 0, 0, 0, 0, 0, time.Now().Location())
			for d := 1; d <= lastDayOfMonth.Day(); d++ {
				dayOfMonth := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Now().Location())
				if d == 13 && dayOfMonth.Weekday() == 5 {
					ret = ret + fmt.Sprintf("%d/%d/%d ", d, m, y)
				}
			}
		}
	}
	return ret
}
