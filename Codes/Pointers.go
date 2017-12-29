package main
import "fmt"

func zeroValue(ival int) {
	ival = 0
}

func zeroPointer(iptr *int) {
	*iptr = 0
}



func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroValue(i)
	fmt.Println("zeroval:", i)
	zeroPointer(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}