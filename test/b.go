package main

import (
	"log"
	"time"
)

func main() {
	currentTime := time.Now()
	log.Println(int(time.Wednesday))
	offset := int(time.Wednesday - 0 + 7)
	weekwendesday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	nextday := weekwendesday.Format("2006-01-02")
	log.Println(nextday)
}
