package main
//Lets explore buffering, synchronization and Direction
/*
Buffering - By default channels are unbuffered, i.e. it only accepts sends (channel <- message) if there is a ready reciever (<- channel).
			Buffered channels however  vaccepted a limited number of messages without having a corresponding reciever for them.

Synchronization - You can use channels to synchronize execution accross GoRoutines.
				  For example you can use a blocking receive to wait for a goroutine to finish.

Direction - When using a channel as a parameter you can specify if the channel is meant to send or recieve values.

 */

import (
	"fmt"
	"time"
)

func slave(done chan bool) {
	fmt.Println("Working ....")
	time.Sleep(time.Second)			// Will wait for a second before sending true into the channel
	fmt.Println("done")

	done <- true
}

func send(ping chan<- string, msg string) {		//Only sends a message
	ping <- msg
}

func accept(ping <- chan string, pong chan<- string) {		//Only recieves a message
	msg := <-ping
	pong <- msg
}

func main()  {

	//Buffering
	message := make (chan int, 2)
	message <- 2		//send values into the channel without a corresponding reciever
	message <- 4
	fmt.Println( <- message, <- message)


	//synchronization
	done := make(chan bool)
	go slave(done)
	<- done



	ping := make(chan string, 1)
	pong := make(chan string, 1)
	send(ping, "Message Sent Sucessfully...")
	accept(ping, pong)
	fmt.Println(<-pong)

}



/*
2 4
Working ....
done
Message Sent Sucessfully...
 */