package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perimeter() float64
}

type rectangle struct{
	width, height float64
}
type circle struct{
	radius float64
}
type square struct{
	side float64
}
type rhombus struct{
	side, dia1, dia2 float64
}

func (r rectangle) area() float64 {
	return r.width*r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, c.radius)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return math.Pow(s.side,s.side)
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

func (rh rhombus) area() float64 {
	return (rh.dia1*rh.dia2)/2
}
func (rh rhombus) perimeter() float64 {
	return 4*rh.side
}



func measure (g geometry) {
	fmt.Println(g)
	fmt.Println("AREA is", g.area())
	fmt.Println("PERIMETER is", g.perimeter())
}

func main(){
	rect := rectangle{7,4}
	circ := circle{7}
	sqr := square{7}
	rhomb := rhombus{13,10,24}


	measure(rect)
	measure(circ)
	measure(sqr)
	measure(rhomb)
}

/*
{7 4}
AREA is 28
PERIMETER is 22
{7}
AREA is 2.587236638715299e+06
PERIMETER is 43.982297150257104
{7}
AREA is 823543
PERIMETER is 28
{13 10 24}
AREA is 120
PERIMETER is 52
 */