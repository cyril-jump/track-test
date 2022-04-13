package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	//check time
	layout := "2006-01-02T15:04:05-07:00"
	now := time.Now()
	//timeStr := now.Format("Mon 02 Jan 2006 3:4:5 MST")
	//timeStr := now.Format("Monday, 02-Jan-06 15:04:05 MST")
	currentTimeStr := "2021-09-19T15:59:41+03:00"

	currentTime, err := time.Parse(layout, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current: ", currentTime)

	fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
	fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))

	fmt.Println(now)
	fmt.Println(now.Truncate(24 * time.Hour))

	//birthday
	HundredYears := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	thisTime := time.Now()
	days := HundredYears.Sub(thisTime).Hours() / 24
	fmt.Println(int64(days))

	/*	//tiker
		timeStart := time.Now()

		for i := 0; i < 10; i++ {
			timer := time.NewTimer(2 * time.Second)
			t := <-timer.C // ожидаем срабатывания таймера
			fmt.Println(t.Sub(timeStart).Seconds())
		}*/

}
