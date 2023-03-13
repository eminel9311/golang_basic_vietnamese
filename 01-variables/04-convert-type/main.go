package main

import (
	"errors"
	"fmt"
	"strconv"
)

func convertStringToFloat(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New("Error: input string is not a valid float")
	}
	return f, nil
}
func main() {

	var a float32 = 10.5
	var b = int(a)
	var c = "22"
	var d, _ = strconv.Atoi(c)
	var e = 33
	var f = strconv.Itoa(e)
	var g = "3.14"
	var h, err = convertStringToFloat(g)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("h have value is %v and have type is %T\n", h, h)
	}

	var i int32 = 10
	var y = float32(i)

	fmt.Printf("a have value is %v and have type is %T\n", a, a)
	fmt.Printf("b have value is %v and have type is %T\n", b, b)
	fmt.Printf("c have value is %v and have type is %T\n", c, c)
	fmt.Printf("d have value is %v and have type is %T\n", d, d)
	fmt.Printf("e have value is %v and have type is %T\n", e, e)
	fmt.Printf("f have value is %v and have type is %T\n", f, f)
	fmt.Printf("g have value is %v and have type is %T\n", g, g)
	fmt.Printf("i have value is %v and have type is %T\n", i, i)
	fmt.Printf("y have value is %v and have type is %T\n", y, y)
}
