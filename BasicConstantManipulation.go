package main

import "fmt"
import "math"

// 'const' declares a constant value.
// GO supports constants of character, string, boolean and numeric values. 
const s string = "constant"

func main() {

fmt.Println(s)
const n = 500000000	 // A 'const' statement can appear anywhere a 'var' statement can.
const m = 30	


const d = 3e20 / n	// Constant expressions perform arithmetic with arbitrary precision.
fmt.Println(d)


fmt.Println(int64(d))  		// A numeric constant has no type until it's given one by an explicit cast.

fmt.Println(math.Sin(m))

// A number can be given a type by using it in a context that requires one such variable assignment or function call.
// Example: 'math.Sin' expects a 'float64'.

}

/*
constant
6e+11
600000000000
-0.9880316240928618
*/

