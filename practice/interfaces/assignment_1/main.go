package main


import (
	"fmt"
)

type shape interface{
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}
type square struct{
	side float64
}


func main(){

	var t triangle
	t.base  = 4
	t.height = 5

	var s square
	s.side = 7

	printArea(t)
	printArea(s)

}


func printArea(s shape){
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64{
	return (.5 * t.base * t.height)
}

func (s square) getArea() float64{
	return (s.side * s.side)
}
