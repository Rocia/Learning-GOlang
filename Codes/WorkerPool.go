package main

import (
	"fmt"
	"time"
)


func worker(id int, jobs <-chan int, results chan<- int) {			//Worker that recieves messages on the jobs channels and sends its responses on the responses channel
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)										//The worker sleeps for 1 second between the reception of a message and sending of a response
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)				//channel to send workers messages
	results := make(chan int, 100)


	for w := 1; w <= 3; w++ {				//Starts 3 workers. However they are blocked as there are no messages in the jobs channel yet.
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {				//Sends 5 messages on the jobs channel
		jobs <- j
	}
	close(jobs)								//Closes the jobs channel to indicate that no more messages will be sent over it

	for a := 1; a <= 5; a++ {				//Collect Results from all the workers
		<-results
	}
}


/*
worker 3 started  job 1
worker 1 started  job 2
worker 2 started  job 3
worker 3 finished job 1
worker 1 finished job 2
worker 2 finished job 3
worker 3 started  job 4
worker 2 started  job 5
worker 3 finished job 4
worker 2 finished job 5
 */
