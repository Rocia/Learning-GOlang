package main
//Lets explore buffering, synchronization and Direction
/*
Buffering - By default channels are unbuffered, i.e. it only accepts sends (channel <- message) if there is a ready reciever (<- channel).
			Buffered channels however  vaccepted a limited number of messages without having a corresponding reciever for them.

Synchronization - You can use channels to synchronize execution accross GoRoutines.
				  For example you can use a blocking receive to wait for a goroutine to finish.

Direction -


 */

import (
	"fmt"
	"time"
)

func slave(done chan bool) {
	fmt.Println("Working ....")
	time.Sleep(time.Minute)			// Will wait for a minute before sending true into the channel
	fmt.Println("done")

	done <- true
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


}



/*
2 4
Working ....
done
 */