package main
//Tickers are used when an action is to be repeated after regular intervals of time.


import (
	"time"
	"fmt"
)

func main() {


	ticker := time.NewTicker(time.Millisecond * 1800)		// ticker for a 500ms interval
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 11600)			//Stops the ticker after 1600 ms
	ticker.Stop()
	fmt.Println("Ticker stopped",time.Now())
}

/*
Tick at 2018-01-11 13:49:12.355650602 +0530 IST
Tick at 2018-01-11 13:49:14.155669228 +0530 IST
Tick at 2018-01-11 13:49:15.955651934 +0530 IST
Tick at 2018-01-11 13:49:17.755644041 +0530 IST
Tick at 2018-01-11 13:49:19.555640916 +0530 IST
Tick at 2018-01-11 13:49:21.355667426 +0530 IST
Ticker stopped 2018-01-11 13:49:22.155672801 +0530 IST
 */