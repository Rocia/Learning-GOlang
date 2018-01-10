package main
// By default Channel send and recieve functions are blocking in nature. However it is possible to implement them as non blocking functions by using the select feature in go.
import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:					// If there is a message in the channel it executes the corresponding case otherwise it executes the default case.
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:				// A non blocking send works by executing the default  case unless the reciever is ready
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:				//We can also implement multiple select cases to have multiple channel actions executed in different circumstances.
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
no message received
no message sent
no activity
 */