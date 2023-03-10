package main

import (
	"basic/variables/02-scope/exportGo"
	"fmt"
)

// biến package
var packageVariable int32 = 10

func main() {
	// biến có phạm vi khối chỉ truy cập được trong hàm main
	var blockVariable int32 = 20
	if true {
		blockVariable2 := "hello everybody"
		fmt.Println(blockVariable, blockVariable2) // có thể truy cập biến trong chính block và block cha
	}

	fmt.Println(packageVariable)           // có thể truy cập biến package ở đây
	fmt.Println(blockVariable)             // có thể truy cập biến block ở đây
	fmt.Println(exportGo.ExportedVariable) // biến từ package khác vẫn có thể truy cập nếu được export
	//fmt.Println(blockVariable2)  // lỗi biên dịch, biến blockVariable2 đã bị hủy

}
