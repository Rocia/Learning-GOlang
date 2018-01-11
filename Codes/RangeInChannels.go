package main

import "fmt"

func main() {
	queue := make (chan string, 2)
	queue <- "First"
	queue <- "Second"

	close(queue)

	for elem := range queue{
		fmt.Println(elem)
	}
}

/*
First
Second
 */