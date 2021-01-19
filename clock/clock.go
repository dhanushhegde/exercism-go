package clock

import (
	"fmt"
	"strconv"
)

type Clock struct {
	h int
	m int
}

func New(h int, m int) Clock {
	return Clock{h: h, m: m}
}

func (c Clock) String() string {
	var hourString string
	var minuteString string
	if hours := (c.h + c.m/60) % 24; hours < 10 {
		hourString = "0" + strconv.Itoa(hours)
	} else {
		hourString = strconv.Itoa(hours)
	}
	if minutes := c.m % 60; minutes < 10 {
		minuteString = "0" + strconv.Itoa(minutes)
	} else {
		minuteString = strconv.Itoa(minutes)
	}

	return fmt.Sprintf("%s:%s", hourString, minuteString)
}

func (c Clock) Add(a int) Clock {
	return Clock{h: c.h + a/60, m: c.m + a%60}
}

func (c Clock) Subtract(a int) Clock {
	return Clock{h: c.h - a/60, m: c.m - a%60}
}
