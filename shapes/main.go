package main

import (
	"fmt"
)

type square struct{
	sideLength float64
}
type triangle struct{
	base float64
	height float64
}

type shape interface{
	getArea() float64
}
func main(){
	sq:=square{4}
	tr:=triangle{3,6}
	printArea(sq)
	printArea(tr)
}
func(sq square) getArea() float64{
	return sq.sideLength*sq.sideLength
}
func(tr triangle) getArea() float64{
	return 0.5*tr.base*tr.height
}
func printArea(s shape){
	fmt.Println(s.getArea())
}