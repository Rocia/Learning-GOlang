// Defer is the python equivalent of finally it is used to ensure that a function is called at the end of execution of a program
package main

import (
	"fmt"
	"os"
)


func main() {

	f := createFile("/tmp/defer.txt")	 	//create a file and write to it and close it with the defer function
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {			//function to create file
	fmt.Println("creating file", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {					//function to write to file
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {					//function to close file
	fmt.Println("closing")
	f.Close()
}



/*
creating file /tmp/defer.txt
writing
closing
 */