package main
//The panic is a method for handling unexpected errors gracefully.
import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")			//Create a temp file
	if err != nil {
		panic(err)
	}
}

/*
panic: a problem

goroutine 1 [running]:
panic(0x47e400, 0xc82000a250)
	/usr/lib/go-1.6/src/runtime/panic.go:481 +0x3e6
main.main()
	/home/tr-dt-102/Desktop/GOlang/Learning-GOlang/Codes/Panic.go:7 +0x65
 */