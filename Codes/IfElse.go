package main

import (
	"fmt"
)

func main(){
	//if else
	if 5%2 == 0 {
		fmt.Println("7 is odd")
	} else {
		fmt.Println("7 is odd")
	}

	//if
	if 8%4 == 0{
		fmt.Println("8 is divisible by 4 ")
	}

	//if - else if - else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}

/*
7 is odd
8 is divisible by 4
9 has 1 digit
*/

