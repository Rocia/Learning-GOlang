package main
import(
	"errors"
	"fmt"
)

func function(arg int)(int,error){
	if arg == 42 {
		return -1, errors.New("I don't like the number 42. Try another")
	}
	return arg+3, nil
}


type argError struct {
	arg int
	prob string
}


func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func function2(arg int)(int,error){
	if arg == 42 {
		return -1, errors.New("I don't like the number 42. Try another")
	}
	return arg+3, nil
}

func main(){
	for _, i := range[]int{7,42} {
		if r, e := function(i);e != nil{
			fmt.Println("Function 1 failed", e)
		} else {
			fmt.Println("Function 2 worked", r)
		}
	}

	for _, i := range[]int{7,42} {
		if r, e := function2(i);e != nil{
			fmt.Println("Function 1 failed", e)
		} else {
			fmt.Println("Function 2 worked", r)
		}
	}
	_, e := function2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}


/*
Function 2 worked 10
Function 1 failed I don't like the number 42. Try another
Function 2 worked 10
Function 1 failed I don't like the number 42. Try another
 */

