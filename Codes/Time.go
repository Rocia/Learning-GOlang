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

	time_formating()

}

func time_formating() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(		// Time parsing uses the same layout values as `Format`.
		time.RFC3339,
		"2018-01-18T22:08:41+05:30")
	p(t1)

	p(t.Format("3:04PM"))			//format and parse work in a manner similar to strftime in python.
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2018-01-18T15:04:05.999999-07:00"))
	form := "3 04 PM"						//Parse time by format
	t2, e := time.Parse(form, "8 41 AM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",	//For pure numeric representations you can use standard formatting.
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"				//Returns errors on malformed data
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}


/*
2018-01-19 15:27:38.371095243 +0530 IST
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
-3h43m20.280291994s
-3.722300081109444
-223.33800486656668
-13400.280291994
-13400280291994
2018-01-19 09:57:38.371095243 +0000 UTC
2018-01-19 17:24:18.931679231 +0000 UTC
2018-01-19T15:27:38+05:30
2018-01-18 22:08:41 +0530 IST
3:27PM
Fri Jan 19 15:27:38 2018
19018-01-18T15:27:38.371227+05:30
0000-01-01 08:41:00 +0000 UTC
2018-01-19T15:27:38-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
 */