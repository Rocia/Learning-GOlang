package main
import "fmt"
func main() {
	s := make([]string,10)
	fmt.Println("Empty Array with spaces as declared:", s)
	s[0], s[4], s[9] = "A", "B", "C"
	fmt.Println("Array after filling index positions 0,4 and 9:", s)
	fmt.Println("Get element at position 4",s[4],"and 5", s[5], ".")
	fmt.Println("Length of array is ", len(s))
	s = append(s, "D")
	s = append(s,"E","F","G")
	fmt.Println("After appension s = ", s)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("After copying array s to c. c = ", c)

	l , m, n := s[2:5], s[:5] , s[4:]
	fmt.Println("Slices l [2:5], m [:5] & n [4:]", l, m, n)

	t := []string {"1","2","3"}
	fmt.Println("Declaration: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array: ", twoD)
}


/*
Empty Array with spaces as declared: [         ]
Array after filling index positions 0,4 and 9: [A    B     C]
Get element at position 4 B and 5  .
Length of array is  10
After appension s =  [A    B     C D E F G]
After copying array s to c. c =  [A    B     C D E F G]
Slices l [2:5], m [:5] & n [4:] [  B] [A    B] [B     C D E F G]
Declaration:  [1 2 3]
2D array:  [[0] [1 2] [2 3 4]]
 */