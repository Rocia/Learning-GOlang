package main
import "fmt"


func Function(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}


func main() {
	Function("direct")
	go Function("GOroutine")
	go func (msg string) {
		fmt.Println(msg)
	} ("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

/*
direct : 0
direct : 1
direct : 2
going
GOroutine : 0
GOroutine : 1
GOroutine : 2
 */