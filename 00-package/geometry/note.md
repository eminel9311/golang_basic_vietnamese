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
4. đây là thứ tự mà Go sử dụng để khai báo và thiết lập các thành phần khác nhau trong một package: import --> const --> var --> init()
    - import: Đầu tiên, chúng ta khai báo các gói (packages) cần thiết cho package hiện tại bằng từ khóa import.
    Các gói này sẽ cung cấp các tính năng mà chúng ta sử dụng trong package của mình.
    - const: Tiếp theo, chúng ta khai báo các hằng số (constants) cho package bằng từ khóa const.
    Các hằng số này sẽ là các giá trị không thay đổi trong suốt thời gian chạy của chương trình.
    - var: Sau đó, chúng ta khai báo các biến (variables) cho package bằng từ khóa var.
    Các biến này có thể thay đổi giá trị trong suốt thời gian chạy của chương trình.
    - init(): Cuối cùng, chúng ta sử dụng hàm init() để khởi tạo các giá trị cho các biến và hằng số trong package. Hàm init() sẽ được gọi trước khi chương trình chính (hàm main()) được thực thi.
    Chúng ta có thể có nhiều hàm init() trong package, và chúng sẽ được thực thi theo thứ tự khai báo.ví dụ chúng ta sử dụng hàm init để valid data của rectLen, rectWidth
```go
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
```