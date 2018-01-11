package main
//Channels may also be terminated after execution to indicate no more values will be sent on that channel.
import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)


	go func() {
		for {
			j, more := <-jobs   //The jobs channel recieves two  messages here, one being the main message and the other a true or false to indicate if more messages are expected over the channel.
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//
	for j := 1; j <= 3; j++ {			//WILL SEND 3 JOBS TO THE CHANNEL
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)						//This sends False as the 'more' argument to the jobs channel if there are no more messages to be sent over the jobs channel.
	fmt.Println("sent all jobs")


	<-done  		//WAIT FOR THE CHANNEL USING SYNC MECHANISM.
}



/*
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
 */