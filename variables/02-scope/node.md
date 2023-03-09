`I.Trong Golang, có 3 phạm vi biến (variable scope) khác nhau, đó là:`
1. Phạm vi biến toàn cục (Global variable scope): Biến được khai báo bên ngoài tất cả các hàm, phương thức và khối lệnh(block).
Biến toàn cục có thể truy cập từ bất kỳ nơi nào trong chương trình.
2. Phạm vi biến cục bộ(Local variable scope): Biến được khai báo trong hàm, phương thức hoặc khối lệnh và chỉ có thể truy cập trong phạm vi của nó.
Khi hàm hoặc phương thức kết thúc, biến cục bộ sẽ bị hủy và không còn tồn tại.
3. Phạm vi biến tham số(Parameter variable scope):  Biến được khai báo trong danh sách tham số của hàm hoặc phương thức và có thể được truy cập và sử dụng trong phạm vi của nó.
ví dụ:
```go
package main

import "fmt"

// Biến toàn cục
var globalVariable = 10

func main() {
// Biến cục bộ
var localVariable = 20

// Biến tham số
funcWithParam(localVariable)

// In giá trị của biến toàn cục và biến cục bộ
fmt.Println("Value of global variable:", globalVariable)
fmt.Println("Value of local variable:", localVariable)
}

// Hàm với tham số
func funcWithParam(paramVariable int) {
// In giá trị của biến tham số
fmt.Println("Value of parameter variable:", paramVariable)
}

```