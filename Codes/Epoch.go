package main

import "fmt"
import "time"

func main() {

	now := time.Now()				// Now and Unix in Time elapse
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000		//since epoch has no milliseconds you will have to divide it to get an answer in milliseconds
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))	//Convert seconds to epoch for the corresponding time
	fmt.Println(time.Unix(0, nanos))
}

/*
2018-01-19 13:51:07.777634576 +0530 IST
1516350067
1516350067777
1516350067777634576
2018-01-19 13:51:07 +0530 IST
2018-01-19 13:51:07.777634576 +0530 IST

 */

