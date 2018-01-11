package main
//We can use a mutex to safely access data across multiple goroutines.
import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)					//Create a dictionary

	var mutex = &sync.Mutex{}						//Declare  a mutex Lock


	var readOps uint64 = 0							//Counters to track read and write ops
	var writeOps uint64 = 0

	for r := 0; r < 100; r++ {						//Start 100 Go routines
		go func() {
			total := 0
			for {

				// For each read we pick a key to access,
				// `Lock()` the `mutex` to ensure
				// exclusive access to the `state`, read
				// the value at the chosen key,
				// `Unlock()` the mutex, and increment
				// the `readOps` count.
				key := rand.Intn(5)				//Acquire a key to lock the mutex variable
				mutex.Lock()						//Lock Variable
				total += state[key]					//read key
				mutex.Unlock()						//Unlock variable
				atomic.AddUint64(&readOps, 1)	//Update atomic counter by 1


				time.Sleep(time.Millisecond)		//Wait for 1ms between read Ops
			}
		}()
	}

	for w := 0; w < 10; w++ {						//start 10 write GO routines
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()						//Acquire Lock
				state[key] = val					//Write op
				mutex.Unlock()						//Release Lock
				atomic.AddUint64(&writeOps, 1)//Increment Write Counter
				time.Sleep(time.Millisecond)		//Sleep Time of 1 ms before starting the next go routine
			}
		}()
	}


	time.Sleep(time.Second)				//Let Both sleep for a second

	readOpsFinal := atomic.LoadUint64(&readOps)		//fetch counter value for read
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)	//fetch counter value for write
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)				//lOCK AND DISPLAY THE STATe dictionary
	mutex.Unlock()
}


/*
readOps: 89976
writeOps: 8999
state: map[1:70 4:80 0:19 2:67 3:50]

 */