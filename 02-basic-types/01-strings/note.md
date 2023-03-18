`Kiểu dữ liệu string trong golang`

1.Trong golang, kiểu dữ liệu string là một kiểu chuỗi các ký tự Unicode.Kiểu dữ liệu này
được sử dụng để biểu diễn các dữ liệu văn bản, như tên, địa chỉ, nội dung của tài liệu văn bản,vv..
để khai báo biến string trong golang như sau
```go
var str string = "Hello World"
```
2.Để truy cập các ký tự trong chuỗi: chúng ta có thể sử dụng chỉ mục của chúng, bắt đầu từ 0.
```go
str := "Hello World"
fmt.Println(str[0])   // In ra ký tự 'H'
fmt.Println(str[1])   // In ra ký tự 'e'
```
3.Golang cung cấp một số phương thức chuỗi cho phép thao tác trên chuỗi,
ví dụ như len(), strings.ToLower(), strings.ToUpper(), strings.TrimSpace(), v.v. Ví dụ:
```go
str := "   Hello World   "
fmt.Println(len(str))                     // In ra độ dài chuỗi (17)
fmt.Println(strings.TrimSpace(str))       // Loại bỏ khoảng trắng thừa ở đầu và cuối chuỗi
fmt.Println(strings.ToLower(str))         // Chuyển đổi chuỗi sang chữ thường
fmt.Println(strings.ToUpper(str))         // Chuyển đổi chuỗi sang chữ hoa
```
4.Nối chuỗi: Chúng ta có thể nối hai chuỗi bằng toán tử + hoặc sử dụng phương thức strings.Join(). Ví dụ:
```go
str1 := "Hello"
str2 := "World"
str3 := str1 + " " + str2               // Nối chuỗi bằng toán tử +
fmt.Println(str3)                       // In ra chuỗi "Hello World"
str4 := []string{"Hello", "World"}
str5 := strings.Join(str4, " ")         // Nối chuỗi bằng phương thức strings.Join()
fmt.Println(str5)                       // In ra chuỗi "Hello World"
```