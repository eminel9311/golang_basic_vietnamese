`I. Shadow trong golang là gì?`
1. Shadow trong golang xảy ra khi một biến được khai báo trong phạm vi của một block(ví dụ là 1 hàm), và tên biến này trùng
với tên của một biến khác được khai báo bên ngoài phạm vị block đó.Khí đó, biến trong phạm vi block bên trong sẽ "che khuất"(shadow)
biến có phạm vi bên ngoài và các thao tác truy cập đến biến này chỉ áp dụng bên trong, khi ra bên ngoài thì biến này sẽ được trả lại giá
trị của nó.
```go
package main

import "fmt"

var x int = 10

func main() {
fmt.Println(x) // In ra giá trị biến x ở phạm vi bên ngoài (10)

if true {
var x int = 20 // Biến x trong phạm vi này "che khuất" biến x ở phạm vi bên ngoài
fmt.Println(x) // In ra giá trị biến x ở phạm vi bên trong (20)
}

fmt.Println(x) // In ra giá trị biến x ở phạm vi bên ngoài (10)
}

```
Khi thực thi, chương trình sẽ in ra các giá trị như sau

```
10
20
10
```
2. Các biến bị shadow có thể dẫn đến sai sót không muong muốn trong quá trình lập trình, do đó, ta nên tránh đặt bến trùng tên với biến khác ở phạm vi bên ngoài