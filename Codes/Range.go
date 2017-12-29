package main
import "fmt"

func main(){
	//Range bin lists
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of nums in array {2,3,4} = ", sum)

	for i, num := range nums {
		fmt.Println("Number", num, "has index:", i)
	}

	// Range in dictionaries
	dictionary := map[string]string{"A": "Apple", "B": "Banana", "C": "Cherry"}
	for k, v := range dictionary {
		fmt.Printf("%s -> %s\n", k, v) //Printf here is used to specify the output format
	}


	for k := range dictionary {			//Prints the keys for the above dictionary
		fmt.Println("Key:", k)
	}

	for i, c := range "ABCDE" {
		fmt.Println("The ascii value of the element at index",i,"is",c)
	}

}
/*
Sum of nums in array {2,3,4} =  9
Number 2 has index: 0
Number 3 has index: 1
Number 4 has index: 2
A -> Apple
B -> Banana
C -> Cherry
Key: A
Key: B
Key: C
The ascii value of the element at index 0 is 65
The ascii value of the element at index 1 is 66
The ascii value of the element at index 2 is 67
The ascii value of the element at index 3 is 68
The ascii value of the element at index 4 is 69
*/
