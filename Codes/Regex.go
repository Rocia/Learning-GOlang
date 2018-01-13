// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

package main

import(
	"bytes"
	"fmt"
	"regexp"
)
var print func(a ...interface{}) (n int, err error) = fmt.Println		//We alias fmt.Println as print

func main() {

	match, _ := regexp.MatchString("p([a-z]+)za", "pizza")	//String match checker
	print(match)

	r, _ := regexp.Compile("p([a-z]+)za")							//REGEX compiler

	print(r.MatchString("pizza"))									//REGEX match

	print(r.FindString("pepperoni pizza"))							//REGEX find

	print(r.FindStringIndex("pepperoni pizza"))						//REGEX find that return start and end indexes of the first occurrence

	print(r.FindStringSubmatch("pepperoni pizza"))					//returns partial and full match details

	print(r.FindStringSubmatchIndex("pepperoni pizza"))				//Index of partial match

	print(r.FindAllString("puffy pezza pepperoni pizza", -1))		//Find All occurrences

	print(r.FindAllStringSubmatchIndex("puffy pezza pepperoni pizza", -1)) //Fins all submatch indexes

	print(r.FindAllString("puffy pezza pepperoni pizza", 2))				//Find 2 of All occurrences

	print(r.Match([]byte("pizza")))										//same aas match string

	r = regexp.MustCompile("p([a-z]+)za")							//Assign REGEX compiler to a constant
	print(r)

	print(r.ReplaceAllString("a pizza", "<food>"))			//REGEX replace all


	in := []byte("a pizza")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)							//Apply functio to text and then replace with REGEX
	print(string(out))
}
/*
true
true
pizza
[10 15]
[pizza iz]
[10 15 11 13]
[pezza pizza]
[[6 11 7 9] [22 27 23 25]]
[pezza pizza]
true
p([a-z]+)za
a <food>
a PIZZA
 */