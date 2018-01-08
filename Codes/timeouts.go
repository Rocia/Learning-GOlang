package main

//Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
//This is implemented in GO with channels and select.

import (
	"time"
	"fmt"
)

func main(){
	channel := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)		//takes 2 seconds
		channel<- "result 1"
	}()
	select {
	case res := <-channel:
		fmt.Println(res)
	case <-time.After(time.Second * 1):		//waits for 1 second
		fmt.Println("timeout 1")
	}

	channel2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)			//takes 2 seconds
		channel2 <- "result 2"
	}()
	select {
	case res := <-channel2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):			//waits 3 seconds
		fmt.Println("timeout 2")
	}

}

/*
timeout 1
result 2
 */

