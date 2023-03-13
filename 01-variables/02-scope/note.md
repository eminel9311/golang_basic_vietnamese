`I.Trong Golang, có 2 phạm vi biến chính(variable scope) khác nhau, đó là:`
1. Phạm vi gói(package scope):
    - Biến được khai báo bên ngoài các khối lệnh trong một gói.
    - Các biến này có thể được truy cập từ bất kỳ đâu trong gói.
    - Để khai báo biến có phạm vi gói, ta sử dụng từ khóa "var".
```go
package main

import "fmt"

var x int = 10 // biến x có phạm vi gói trong package main

func main() {
    fmt.Println(x) // có thể truy cập biến x tại đây
}
```
2. Phạm vi khối(block scope):
    - Biến được khai báo trong một khối lệnh nhất định, ví dụ như trong 1 hàm hoặc 1 vòng lặp.
    - Các biến này chỉ được có thể được truy cập trong phạm vi khối lệnh đó và các khối lệnh con của nó.
    - Khi khối lệnh kết thúc, biến sẽ bị hủy và không thể truy cập được nữa.
ví dụ:
```go
func example() {
    x := 10 // biến x có phạm vi khối trong hàm example
    if true {
        y := 20 // biến y có phạm vi khối trong khối if
        fmt.Println(x, y) // có thể truy cập được biến x và y tại đây
    }
    fmt.Println(x) // chỉ có thể truy cập được biến x tại đây
    fmt.Println(y) // lỗi biên dịch, biến y đã bị hủy
}
```

`II. Quy tắc đặt tên biến trong Golang`
1. Tên biến nên bắt đầu bằng một chữ cái hoặc ký tự gạch dưới
2. Tên biến bên sử dụng chữ cái hoặc số và không sử dụng dấu cách hoặc các ký tự đặc biệt
3. Tên biến nên sử dụng kiểu chữ thường cho các biến có phạm vi(scope) trong cùng package,
và sử dụng kiểu chữ hia cho các biến được xuất(exported) ra từ package để sử dụng cho package khác.

```go
package mypackage

var myVariable int = 10 // Biến có phạm vi trong cùng package
var ExportedVariable int = 20 // Biến được xuất ra từ package

```
import và sử dụng các biến này cho package khác

```go
package myotherpackage

import "path/to/mypackage"

func myFunc() {
    // Truy cập biến ExportedVariable trong package mypackage
    myVar := mypackage.ExportedVariable
    fmt.Println(myVar)
}

```