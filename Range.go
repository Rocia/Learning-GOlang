package main
import "fmt"

func main(){
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}


	kvs := map[string]string{"A": "Apple", "B": "Banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) //Printf used to specify output format
	}


	for k := range kvs {
		fmt.Println("key:", k)
	}


	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
/*
Sum: 9
index: 1
A -> Apple
B -> Banana
key: A
key: B
0 103
1 111
*/
