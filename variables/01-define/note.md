`I. Khai báo biến trong golang có thể sử dụng các cách sau`
1. khai báo biến và kiểu dữ liệu
```go
var i int32
i = 10
```
2. khai báo biến và gán dữ liệu luôn cho biến
```go
var j int32 = 11
```
3. khai báo biến sử dụng cú pháp ngắn gọn, khi sử dụng dấu chấm động ":" golang sẽ định nghĩa kiểu dữ liệu dựa vào dữ liệu truyền vào.
```go
k := 12
```
4. khai báo nhiều biến cùng kiểu dữ liệu trong 1 lần
```go
var b, c int = 1, 2
c, d := 1, 2
```
5. khai báo nhiều biến khác kiểu dữ liệu trong 1 lần
```go
var (
n int32  = 100
m string = "xin chao"
h bool   = true
)
```


