# 1. go 标准库之 io.Copy 和 ioutil.ReadAll

## 1.1 介绍
go 标准库中通过 `ioutil.ReadAll` 实现数据流的读取，`io.Copy` 实现数据流的读取和写入。

那两者有什么区别呢？  

有。

ioutil.ReadAll 通过 slice 将数据流读到内存中。slice 会动态扩容，对于大文件的读取会导致内存不足，被 OOM kill。并且频繁的动态扩容也会导致时间消耗增加。  

io.Copy 构造了三种场景：  
1. 如果 Reader 接口类型 src 实现了 WriteTo 方法，则调用 `src.WriteTo(dst)` 将 src 的数据流写到 dst 中。  
2. 如果 Writer 接口类型 dst 实现了 ReadFrom 方法，则调用 `dst.ReadFrom(src)` 将 Reader 接口类型的数据流 src 读到 dst 中。  
3. 如果前两种场景都不满足，则创建固定大小的 slice，将数据从 src 读到 slice，再由 slice 写到 dst 中。因为边读边写，可以实现 slice 大小的固定，不需要动态扩容。  

下面再细化看具体流程。

## 1.2 slice append
go 中 slice 通过 append 函数实现动态扩容：  
- 对于不足 1024bytes 的切片，扩容后的容量为原容量的 2 倍。以防扩容太过频繁。   
- 对于超过 1024bytes 的切皮，扩容后的容量为原容量的 1.25 倍。以防内存占用过多。  

代码示例如下。  
```
s := []byte{104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100}
fmt.Println(s, len(s), cap(s))
s = append(s, 0)[:len(s)]
fmt.Println(s, len(s), cap(s))
```

result:
```
[104 101 108 108 111 44 32 119 111 114 108 100] 12 12
[104 101 108 108 111 44 32 119 111 114 108 100] 12 24
```

关于 slice 的扩容机制实现可参考 [Go 设计与实现：切片](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/)

## 1.3 io.ReadAll 和 io.Copy 实现
### 1.3.1 io.ReadAll
```
func ReadAll(r Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == EOF {
				err = nil
			}
			return b, err
		}
	}
}
```

从代码可以看出，函数 `ReadAll` 首先创建了 512 bytes 的切片 b。接着循环读取数据流到 b。其中，当 b 切片满的时候，调用 append 为 b 扩容，再继续读取。

可以想见这种机制，当读取大文件时频繁的扩容，大文件写入到内存中会导致时间消耗，更甚者导致程序被 OOM kill 掉。

构造一个大文件场景验证猜测：
```
resp, _ := http.Get("xxx")

data, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatalln(err)
}

if err := ioutil.WriteFile("./log.tar", data, 0755); err != nil {
    log.Fatalln(err)
}
```

这里 http.Get 的文件大小为 6.29G，通过 io.ReadAll 将 6.29G 的文件写入到内存(b) 中，接着将写入的数据 data(b) 写到文件 log.tar。
执行程序发现，运行一段时间后，程序报错：
```
signal: killed
```

程序被 kill 掉了，且程序不是正常退出。原因是大文件的内存写入消耗了内存空间，导致 OOM 将该程序 kill 掉。

### 1.3.2 io.Copy
```
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
    // 场景 1
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}

    // 场景 2
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

    // 场景 1，2 都不满足情况下，实现场景 3
    // 首先创建 32KB 的 slice buf 
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}

    // 循环读取和写入数据流
	for {
        // 调用接口类型 src 的 Read 方法将数据流读取到 buf 中
		nr, er := src.Read(buf)
		if nr > 0 {
            // 如果有数据流读取，则调用接口类型 dst 的 Write 方法将 buf 数据写入到 dst 中 
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
```

这里要注意的是，场景 3 是边读边写从而保证 buf 的 size 固定。  

可以想见，对于场景 3，不需要动态扩容，且每次写入 32KB 的数据到内存能保证大文件的批量读写。

继续验证猜测：
```
resp, _ := http.Get("xxx")
file, _ := os.Create("log.tar")

if _, err := io.Copy(file, resp.Body); err != nil {
    log.Fatalln(err)
}
```

程序执行成功，输出文件 log.tar。

## 1.4 Benchmark
通过性能测试看 `io.Copy` 和 `io.ReadAll` 测试性能：

```
func BenchmarkGetFileFromReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFileFromReadAll()
	}
}

func BenchmarkGetFileFromCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFileFromReadAll()
	}
}
```

result:
```
# go test -bench=. -v -benchmem -benchtime=5s -count=3 .
goos: linux
goarch: amd64
pkg: chapter1/1.5/1.7/poc2
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkGetFileFromReadAll
BenchmarkGetFileFromReadAll-3                 30         229097198 ns/op        160770630 B/op       969 allocs/op
BenchmarkGetFileFromReadAll-3                 30         218096082 ns/op        160776824 B/op      1061 allocs/op
BenchmarkGetFileFromReadAll-3                 26         204630763 ns/op        160773033 B/op      1066 allocs/op
BenchmarkGetFileFromCopy
BenchmarkGetFileFromCopy-3                    27         219585784 ns/op        160779221 B/op      1092 allocs/op
BenchmarkGetFileFromCopy-3                    37         213396925 ns/op        160772424 B/op      1062 allocs/op
BenchmarkGetFileFromCopy-3                    28         212014986 ns/op        160772857 B/op      1063 allocs/op
PASS
ok      chapter1/1.5/1.7/poc2   63.964s
```

这里测试的文件大小为 26M，从性能测试上看差异并不大。  
这个结果可能是多方面的，这里不详细展开。要知道的是，实际使用中如果不清楚数据流大小的情况下最好使用 io.Copy。
