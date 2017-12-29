package main
//structures work in a way that is quite similar 5to the inbuilt 'dictionary' datatype in python
//i.e. entities are stored within a structure that is immutable as key value pairs
import "fmt"

type menu struct{
	Dish string
	RateInDollars int
}


func main(){
	fmt.Println(menu{"ravioli", 50})
	fmt.Println(menu{Dish:"Gnocchi", RateInDollars:45})
	fmt.Println(menu{Dish:"Lasagne"}) // 0 int for RateInDollars
	fmt.Println(menu{RateInDollars:55}) //Blank string for Dish
	fmt.Println(&menu{Dish:"Spaghetti", RateInDollars:40}) //& Prefix will yield a 
}