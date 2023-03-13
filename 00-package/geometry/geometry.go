package main

import (
	"basic/00-package/geometry/rectangle"
	"fmt"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("perimeter of rectangle %.2f\n", rectangle.Perimeter(rectLen, rectWidth))
	fmt.Printf("diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))

}
