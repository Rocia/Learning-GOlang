package main

import (
	"strconv"		//For number parsing
	"fmt"
	"go/constant"
)
func number_parser(){
	f, _ := strconv.ParseFloat("1.234", 64)			//Parse float 64 says how many bits of precision to parse
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)		//Parse integer with base 0 and fit in 64 bits
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)	//Can infer hexadecimal numbers as well
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)		//Parse u int
	fmt.Println(u)

	k, _ := strconv.Atoi("135")								//'Atoi' convenient keyword for parsing integers with base 10
	fmt.Println(k)

	_, e := strconv.Atoi("What")								//return errors and bad input
	fmt.Println(e)
}


func main() {
	number_parser()

}