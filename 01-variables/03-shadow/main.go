package main

import "fmt"

var x int32 = 10

func main() {
	fmt.Println(x) // In ra giá trị biến x ở phạm vi bên ngoài (10)

	if true {
		var x int32 = 20 // Biến x trong phạm vi này "che khuất" biến x ở phạm vi bên ngoài
		fmt.Println(x)   // In ra giá trị biến x ở phạm vi bên trong (20)
	}
	fmt.Println(x) // In ra giá trị biến x ở phạm vi bên ngoài (10)

}
