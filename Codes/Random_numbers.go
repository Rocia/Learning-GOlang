package main

import "time"
import "fmt"
import "math/rand"

func main() {

	fmt.Print(rand.Intn(100), ",",rand.Intn(100))		//randint below 100
	fmt.Println()

	fmt.Println(rand.Float64())				// float 0.0 <= f < 1.0

	fmt.Print((rand.Float64()*5)+5, ",", (rand.Float64()*5)+5) 		//5.0 <= f' < 10.0`.
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())		//the default number generator is deterministic nd thus produces the same sequence everytime unless you give it a seed
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",", r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)	// If you seed a source with the same number, it produces the same sequence of random numbers.
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",", r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",",r3.Intn(100))
	fmt.Print()
}

/*
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
89,25
5,87
5,87
 */