package main

import (
	"basic/00-package/geometry/rectangle"
	"fmt"
	"log"
)

var rectLen, rectWidth float64 = 6, 7

/*
 * Hàm init được thêm vào
 */
func init() {
	fmt.Println("Rectangle package initialized")
	if rectLen < 0 {
		log.Fatal("Length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("Width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("perimeter of rectangle %.2f\n", rectangle.Perimeter(rectLen, rectWidth))
	fmt.Printf("diagonal of rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))

}
