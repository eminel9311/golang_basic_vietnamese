package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   Hello World   "
	fmt.Println(len(str))               // In ra độ dài chuỗi (17)
	fmt.Println(strings.TrimSpace(str)) // Loại bỏ khoảng trắng thừa ở đầu và cuối chuỗi
	fmt.Println(strings.ToLower(str))   // Chuyển đổi chuỗi sang chữ thường
	fmt.Println(strings.ToUpper(str))   // Chuyển đổi chuỗi sang chữ hoa

	str1 := "Hello"
	str2 := "World"
	str3 := str1 + " " + str2 // Nối chuỗi bằng toán tử +
	fmt.Println(str3)         // In ra chuỗi "Hello World"
	str4 := []string{"Hello", "World"}
	str5 := strings.Join(str4, " ") // Nối chuỗi bằng phương thức strings.Join()
	fmt.Println(str5)               // In ra chuỗi "Hello World"

}
