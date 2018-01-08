package main

import (
	"time"
	"fmt"
)

//The select feature in GO lets you wait on multiple channel operations. Combining goroutines and channels with select makes Go a Powerful language.


func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		channel1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		channel2 <- "two"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		}
	}
}

/*
received one
received two
 */