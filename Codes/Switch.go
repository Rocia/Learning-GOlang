package main
import (
	"fmt"
	"time"
)

func main()  {
	i := 2
	fmt.Print("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// switch without any expression works in the same way an if else would
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("It's",t.Hour(), "in the morning")
	default:
		fmt.Println("It isn't morning")
	}
	variable := func(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("Its a BOOLEAN")
	case int:
		fmt.Println("Its an INTEGER")
	default:
		fmt.Println( t, "is of unknown type %T\n")
	}
	}
	variable(true)
	variable(1)
	variable("hey")

}


