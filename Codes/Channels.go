package main

//Pipes that connect concurrent goRoutines

//Used to send and recieve values between goRoutines
import "fmt"


func add(a int, b int) (int, int) {
	return a + b, b
}

func mult(c int, d int) int {
	return c * d
}


func main() {

	//Example 1

	messages := make(chan string)		//make a channel typed with the value it passes

	go func() { messages <- "call" }()		// Now send a message into that channel by calling it in a goRoutine

	msg := <-messages		//Fetch the value from the channel and print it
	fmt.Println(msg)		//By default "block" is sent and recieved till both functions are ready



	//Example 2

	number := make (chan int, 2)		// make a channel of size 2

	num1, num2 := add (3, 4)		// take the return from add function and store within the channel
	number <- num1
	number <- num2

	res := mult(<-number, <-number) 	// assign those values to the mult function
	fmt.Println(res)


}


/*
call
28
 */