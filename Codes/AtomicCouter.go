package main
// atomic counters are among the state management mechanisms in Go. it is used to manage state among multiple Go routines
import (
	"fmt"
 	"time"
 	"sync/atomic"
)

func main() {

	var counter uint64 = 0		//We use the unsigned integer datatype to ensure that our value is always positive.

	for i := 0; i < 50; i++ {			//To show concurrent updates to the counter we start 50 go routines
		go func() {
			for {

				atomic.AddUint64(&counter, 1)			//This automatically increments the counter because of the AddUint64 which gives it the memory address of our 'counter' with the & syntax

				time.Sleep(time.Millisecond*10)				//Wait for 2 ms between increments
			}
		}()
	}

	time.Sleep(time.Second)									//Sleeps for one second to allow some valuesto accumulate in the counter


	counterFinal := atomic.LoadUint64(&counter)					//To safely use the counter we extract a copy of it to CounterFinal from its memory address given by &counter
	fmt.Println("ops:", counterFinal)

	}



/*
ops: 4950
 */