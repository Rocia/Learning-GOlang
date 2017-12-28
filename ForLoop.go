package main
import "fmt"
func main(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	
	for {
		fmt.Println("Loop") // Loop will run once for break and infinitely for continue.
		break
	}
	
	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

/*
1
2
3
7
8
9
Loop
1
3
5
*/

