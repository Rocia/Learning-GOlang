package main

import "fmt"

//There are two main kinds of functions we will explore here
//Anonymous and Recursive

func RescursiveFactorial(a int) int {						//Functions called by any number of trailing arguments
	if a == 0 {
		return 1
	} else {
		return a * RescursiveFactorial(a-1)
	}
}



func AnonymousFibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func iterations() int {
	return 5
}

func main(){
	Anonymous := AnonymousFibonacci()
	iterations := iterations()
	fmt.Println("Recursive Factorial", RescursiveFactorial(iterations))
	fmt.Println("Anonymous Fibonaccci")
	for iterations != 0{
		fmt.Println(Anonymous())
		iterations --
	}
}

/*
Recursive Factorial 120
Anonymous Fibonaccci
1
1
2
3
5
 */