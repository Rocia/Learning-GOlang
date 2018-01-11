package main
//Rate limiting is used to restrict the resource utilization and to Ensure the quality of services.

import (
	"time"
	"fmt"
	)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {				//Limit the requests to the channel named 'requests'
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)			//works as a regulator ande recieves a value every 200ms


	for req := range requests {								//We block our requests till we get a response the 'limiter channel which in turn serves as a 200ms delay.
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)				//Allows a burst of upto 3 events every 200 ms


	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()							// Fill up the channel to view allowed bursting.
	}


	go func() {
		for t := range time.Tick(time.Millisecond * 1800) {	//Every 200 ms we add a value to the burstylimiter channel
			burstyLimiter <- t
		}
	}()


	burstyRequests := make(chan int, 5)						//Send 5 requests of which only 3 must be displayed
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
request 1 2018-01-11 15:01:32.186046339 +0530 IST
request 2 2018-01-11 15:01:32.386040709 +0530 IST
request 3 2018-01-11 15:01:32.586052844 +0530 IST
request 4 2018-01-11 15:01:32.786073472 +0530 IST
request 5 2018-01-11 15:01:32.986052525 +0530 IST
request 1 2018-01-11 15:01:32.986090362 +0530 IST
request 2 2018-01-11 15:01:32.986093998 +0530 IST
request 3 2018-01-11 15:01:32.986095728 +0530 IST
request 4 2018-01-11 15:01:34.786163278 +0530 IST
request 5 2018-01-11 15:01:36.586181911 +0530 IST
 */