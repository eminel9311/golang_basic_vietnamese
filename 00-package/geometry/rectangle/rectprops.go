package rectangle

import "math"

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Perimeter(len, wid float64) float64 {
	perimeter := (len + wid) * 2
	return perimeter
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
