# Note

## 变量

变量的四种声明方式:
```
s := ""
var s string
var s = ""
var s string = ""
``` 

其中，常用的是短变量声明 `s := ""` 和 `var s string`。  
第三种形式在声明多变量时会用到，很少用。第四种声明增加了冗余，不够简洁。  

第一种短变量声明简洁，不过只能用在函数体内。第二种依赖默认初始化。实际使用时，使用隐式的初始化来表明初始化变量不重要，使用显示初始化表明初始化变量的重要性。

## 基准测试

基准测试在测试之前要想好测试场景和测试用例，方便构造测试函数。  


实践中，使用 `benchmark` 函数封装一层的方式更为优雅。如下：
```
func BenchmarkPlusContact(b *testing.B)           { benchmark(b, plusContact) }
func BenchmarkStringsJoinContactPre(b *testing.B) { benchmark(b, stringsJoinContactPre) }
func BenchmarkStringsJoinContact(b *testing.B)    { benchmark(b, stringsJoinContact) }
```

若不加封装示例如下：
```
func BenchmarkPlusContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusContact(30, randomString(10))
	}
}

func BenchmarkStringsJoinContactPre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoinContactPre(30, randomString(10))
	}
}
```

两段基准测试函数只在函数调用处不一样。因此，可以通过封装调用函数的方式减少代码函数，使程序更优雅。


执行基准测试，查看输出：
```
# go test -bench=. -v -benchmem -benchtime=5s -count=2 .
goos: linux
goarch: amd64
pkg: chapter1/1.3
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkPlusContact
BenchmarkPlusContact-3                   1000000              6666 ns/op            5368 B/op         31 allocs/op
BenchmarkPlusContact-3                    951889              5597 ns/op            5368 B/op         31 allocs/op
BenchmarkStringsJoinContactPre
BenchmarkStringsJoinContactPre-3         2134222              2966 ns/op            1920 B/op          5 allocs/op
BenchmarkStringsJoinContactPre-3         2448580              2807 ns/op            1920 B/op          5 allocs/op
BenchmarkStringsJoinContact
BenchmarkStringsJoinContact-3            2504148              2138 ns/op            1392 B/op          9 allocs/op
BenchmarkStringsJoinContact-3            2189086              2394 ns/op            1392 B/op          9 allocs/op
PASS
ok      chapter1/1.3    50.663s
```

从输出可以看出，`StringsJoinContact` 相比于 `StringsJoinContactPre` 分配内存次数较多，但分配内存空间较少。


# 引用
- [benchmark 基准测试](https://geektutu.com/post/hpg-benchmark.html)
