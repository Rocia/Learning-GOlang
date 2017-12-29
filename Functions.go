package main
import "fmt"


func add2(a int, b int) int { 			//Adition of two integers
	return a + b		//Like in python GO too requires explicit returns from functions
}

func add3(a, b, c int) int {				//Sets a , b and c to type int
	return a + b + c
}

func AddMult(a,b int) (int, int) {			//Multiple returns from a function
	return a+b, a*b
}
func main() {

	result := add2(4, 3)			//var result will accept the return from the function being called
	fmt.Println("4 + 3 =", result)

	result = add3(1, 2, 4)
	fmt.Println("4 + 2 + 1 =", result)

	result1, result2 := AddMult(4, 3)
	fmt.Println("4 + 3 =", result1,"\n4 * 3 =", result2)
}
