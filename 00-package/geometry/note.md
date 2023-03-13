`Package trong golang là gì?`
1. package trong Golang là một cơ chế nhóm các tệp mã lại với nhau để cung cấp tính tái sử dụng và quản lý
phạm vi các biến, hàm, cấu trúc dữ liệu và các đối tượng khác.Mỗi tệp Go được đặt trong một package nhất định,
và một package có thể được sử dụng bởi các package khác thông qua khai báo import

2. trong 1 ứng dụng golang muốn chương trình thực thi cần có một main function(hàm chính).Hàm này là điểm bắt
đầu(entry point) để thực thi.Main function phải nằm trong main package(gói chính).

3. trong Golang có rất nhiều package chuẩn cung cấp sẵn như "fmt" để định dạng và hiển thị các chuỗi, "io" để xử
lts I/O, "net" để làm việc với mạng, "http" để xây dựng ứng dụng web, vv.Ngoài ra, chúng ta cũng có thể tạo các
package riêng để sử dụng lại trong các dự án khác

```go
//rectprops.go
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


//geometry.go
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


```