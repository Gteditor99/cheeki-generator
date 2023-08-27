package main

import (
	"fmt"
	"time"
)

func kst(t time.Time) string {
    loc, err := time.LoadLocation("Asia/Seoul")
    if err != nil {
        panic("Timezone Asia/Seoul not found.")
    }
    kstTime := t.In(loc)
    formatted := kstTime.Format("2006-01-02")
    return formatted
}

var startDate = kst(time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC))

type classDay struct {
    period    int
    hasKorean bool
    date      time.Time
    weekday   time.Weekday
}

// Korean classes are on Tuesdays (6th period), Wednesdays (6th period), and Fridays (1st period)
func isKoreanClass(c *classDay) bool {
    switch c.date.Weekday() {
    case time.Tuesday:
        c.period = 6
        c.weekday = time.Tuesday
        return true
    case time.Wednesday:
        c.period = 6
        c.weekday = time.Wednesday
        return true
    case time.Friday:
        c.period = 1
        c.weekday = time.Friday
        return true
    default:
        return false
    }
}

func main() {
    startDate, err := time.Parse("2006-01-02", startDate)
    if err != nil {
        panic(err)
    }
    fmt.Println("Start date:", kst(startDate))
    fmt.Println("Today:", kst(time.Now()))
    fmt.Println("Printing every instance of Korean class up to today...")
    for d := startDate; d.Before(time.Now()); d = d.AddDate(0, 0, 1) {
        c := classDay{date: d}
        if isKoreanClass(&c) {
            if c.period != 0 {
                fmt.Println(kst(c.date), c.weekday, "period", c.period)
            } else {
                fmt.Println(kst(c.date), c.weekday)
            }
        }
    }
}