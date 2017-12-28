package main
import "fmt"
func main(){

//strings
var a string = "initial" //explicit declaration
var b = "initial2" //simple declaration
fmt.Println(a)
fmt.Println(b)

//Integers
var c,d int = 12,15
var e int 		//auto set to 0
fmt.Println(c,d)
fmt.Println(e)

//Boolean
var f = false
fmt.Println(f)
fmt.Println(!f)

//Auto type declaration
g := "short"
h := 2
fmt.Println(g)
fmt.Println(h)
}




/*
initial
initial2
12 15
0
false
true
short
2
*/
