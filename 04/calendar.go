package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"time"
)

type cal struct {
	day     int
	weekday time.Weekday
	isToday bool
}

type cals struct {
	year  int
	month string
	cal   []cal
}

func Calendar() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	cals := cals{
		year:  currentYear,
		month: currentMonth.String(),
	}
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	_, _, today := now.Date()
	_, _, lastDayOfMonth := lastOfMonth.Date()
	for i := 0; i < lastDayOfMonth; i++ {
		isToday := false
		if i == today-1 {
			isToday = true
		}
		nextDay := firstOfMonth.AddDate(0, 0, i)
		cals.cal = append(cals.cal, addDay(nextDay, isToday))
	}
	cals.Fprint()
}

func addDay(d time.Time, isToday bool) cal {
	c := cal{}
	_, _, day := d.Date()
	c.day = day
	c.weekday = d.Weekday()
	c.isToday = isToday
	return c
}

func (c cals) Fprint() {
	fmt.Printf("   %s  %d \n", c.month, c.year)
	weekday := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	for _, day := range weekday {
		fmt.Printf("%s ", day)
	}
	fmt.Printf("\n")

	space := int(c.cal[0].weekday)
	for i := 0; i < space; i++ {
		fmt.Printf("   ")
	}

	for _, day := range c.cal {
		add := " "
		if day.day < 10 {
			add = "  "
		}
		if (int(day.weekday)+1)%7 == 0 {
			add = "\n"
		}
		for num, _ := range weekday {
			if int(day.weekday) == num {
				if day.isToday {
					fmt.Printf("%d%s", Gray(24-1, day.day).BgGray(1-1), add)
				} else {
					fmt.Printf("%d%s", day.day, add)
				}
			}
		}
	}
	fmt.Printf("\n")
}
