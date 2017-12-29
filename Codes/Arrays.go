package main
import "fmt"
func main() {
	var a[5] int
	fmt.Println("Empty array:", a)
	a[4] = 50
	fmt.Println("Set variable at index position 4:", a)
	fmt.Println("Get variable at index position 4:", a[4])
	fmt.Println("Length of array:", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("One line declaration of array b as [5]int{1,2,3,4,5}:", b)

	var twoD [2][3] int
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array: ", twoD)
}

/*
Empty array: [0 0 0 0 0]
Set variable at index position 4: [0 0 0 0 50]
Get variable at index position 4: 50
Length of array: 5
One line declaration of array b as [5]int{1,2,3,4,5}: [1 2 3 4 5]
2D array:  [[0 1 2] [1 2 3]]
 */