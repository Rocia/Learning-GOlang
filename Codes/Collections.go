package main

import "strings"
import "fmt"

// Returns the first index of the target string `t`, or
// -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		fmt.Println(i, v)
		if v == t {
			return i
		}
	}
	return -1
}


func Include(vs []string, t string) bool {				//Returns true or false based on wjhether the target stringis a part of the main string
	return Index(vs, t) >= 0
}


func Any(vs []string, f func(string) bool) bool {		//Returns true if any one of the target string returns satisfies the condition dictated by the function
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}


func All(vs []string, f func(string) bool) bool {		//Returns true if all of the target string returns satisfies the condition dictated by the function
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"ravioli", "spaghetti", "carbonara", "farfalle"}

	fmt.Println("Index function for the word 'carbonara'", Index(strs, "carbonara"))

	fmt.Println("Include function for the word 'spaghetti'", Include(strs, "spaghetti"))

	fmt.Println("Any for the letter 'r'",Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "r")
	}))

	fmt.Println("All for the letter 'a'",All(strs, func(v string) bool {
		return strings.HasPrefix(v, "a")
	}))

	fmt.Println("Filter strings",Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println("convert to upper",Map(strs, strings.ToUpper))				//Using named instead of anonymous functions

}

/*
0 ravioli
1 spaghetti
2 carbonara
Index function for the word 'carbonara' 2
0 ravioli
1 spaghetti
Include function for the word 'spaghetti' true
Any for the letter 'r' true
All for the letter 'a' false
Filter strings [spaghetti farfalle]
convert to upper [RAVIOLI SPAGHETTI CARBONARA FARFALLE]

 */