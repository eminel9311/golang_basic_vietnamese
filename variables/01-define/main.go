package main

import "fmt"

func main() {
	// khai báo biến i kiểu interger có kích thước 32bit
	var i int32
	i = 10

	// khai báo biến j kiểu interger và gán giá trị cho biến j có kích thước 32bit
	var j int32 = 11

	// khai báo biến ngắn gọn
	k := 12.5

	// khai báo nhiều biến cùng 1 kiểu dữ liệu trong 1 lần
	var a, b int32 = 1, 2

	// khai báo nhiều biến khác kiểu dữ liệu trong 1 lần
	var (
		n int32  = 100
		m string = "xin chao"
		h bool   = true
	)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("k has type %T and value %v\n", k, k)
	fmt.Println(a, b)
	fmt.Println(n, m, h)
}
