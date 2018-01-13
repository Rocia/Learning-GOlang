package main
//SOrting using built-in and custom methods

import (
	"fmt"
	"sort"
)

type ByLength []string

func custom() []string {
	food := []string{"ravioli", "spaghetti", "carbonara", "farfalle"}
	sort.Sort(ByLength(food))
	return food
}

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	strings := []string{"c", "a", "b"}					//Sorted strings
	sort.Strings(strings)
	fmt.Println("Strings:", strings)

	ints := []int{7, 2, 4}								//Sorted integers
	sort.Ints(ints)
	fmt.Println("Integers:", ints)

	fmt.Println("Sorted:", sort.IntsAreSorted(ints))//check if array is sorted
	
	fmt.Println(custom())								//call the custom function where the
}

/*
Strings: [a b c]
Integers: [2 4 7]
Sorted: true
[ravioli farfalle spaghetti carbonara]
 */