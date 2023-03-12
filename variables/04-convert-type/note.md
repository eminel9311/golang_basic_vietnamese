`Trong golang để convert kiểu dữ liệu chúng ta sử dụng các toán tử hoặc các hàm được cung cấp sẵn`
1. chuyển đổi kiểu int sang kiểu float
```go
var i int32 = 10
var y = float32(i)
```

2. chuyển đồi kiểu float sang kiểu int
```go
var a float32 = 10.5
var b = int(a)
```

3. chuyển đổi kiểu string sang kiểu int chúng ta import thư viện strconv và sử dụng hàm Atoi
```go
var c = "22"
var d, _ = strconv.Atoi(c)
```
bình thường _ sẽ là hàm hanlde err nếu chuỗi không phải số hợp lệ tuy nhiên trong ví dụ này chúng ta không sử dụng nó nên sẽ tạm thời ẩn đi

4. chuyển đổi kiểu số sang string chúng ta cũng import thư viện strconv và sử dụng hàm Itoa
```go
var e = 33
var f = strconv.Itoa(e)
```
5. chuyển đổi kiểu string sang float chúng ta inport thư viện strconv và sử dụng hàm ParseFloat, tuy nhiên trong trường hợp nếu string chúng ta truyền vào
sai định dạng(không phải là 1 số float hợp lệ) thì sẽ không convert được, vì vậy chúng ta sẽ cần handle trường hợp lỗi này có thể viết thêm 1 hàm để valid
trường hợp này như ví dụ dưới đây.
```go
func convertStringToFloat(str string) (float64, error) {
f, err := strconv.ParseFloat(str, 64)
if err != nil {
return 0, errors.New("Error: input string is not a valid float")
}
return f, nil
}

var g = "3.14"
var h, err = convertStringToFloat(g)
if err != nil {
fmt.Println(err)
} else {
fmt.Printf("h have value is %v and have type is %T\n", h, h)
}
```