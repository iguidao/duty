package main

import (
	"log"
	"time"
)

func ExistHoliday(nextday string) bool {
	var holiday [2]string
	holiday[0] = "2020-03-26"
	for _, aaa := range holiday {
		if nextday == aaa {
			return true
		}
	}
	return false
}

func main() {
	nextday := "2020-03-25"
	var b int = 7
	for i := 0; i < b; i++ {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		tt, _ := time.ParseInLocation("2006-01-02", nextday, loc)
		date_day := time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, i)
		nextday := date_day.Format("2006-01-02")
		if !ExistHoliday(nextday) {
			log.Println(nextday)
		} else {
			b = b + 1
		}

	}

}
