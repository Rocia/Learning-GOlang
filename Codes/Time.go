package main

import (
	"fmt"
	"time"
)


func main() {
	p := fmt.Println

	now := time.Now()		//Find time
	p(now)

	mytime := time.Date(					//Time is associated with location
		2018, 01, 19, 13, 40, 58, 651387237, time.UTC)
	p(mytime)

	p(mytime.Year())			//extract different values
	p(mytime.Month())
	p(mytime.Day())
	p(mytime.Hour())
	p(mytime.Minute())
	p(mytime.Second())
	p(mytime.Nanosecond())
	p(mytime.Location())

	p(mytime.Weekday())			//day of the week

	p(mytime.Before(now))		//compares time
	p(mytime.After(now))
	p(mytime.Equal(now))

	diff := now.Sub(mytime)		//Time Interval
	p(diff)

	p(diff.Hours())				//Time Interval in different units
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(mytime.Add(diff))			//+ve and -ve time difference
	p(mytime.Add(-diff))
}



/*
2018-01-19 13:46:16.984418353 +0530 IST
2018-01-19 13:40:58.651387237 +0000 UTC
2018
January
19
13
40
58
651387237
UTC
Friday
false
true
false
-5h24m41.666968884s
-5.411574158023333
-324.6944494814
-19481.666968884
-19481666968884
2018-01-19 08:16:16.984418353 +0000 UTC
2018-01-19 19:05:40.318356121 +0000 UTC
 */