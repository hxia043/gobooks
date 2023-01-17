# 前言
基本数据包括的知识点不少，平时可能都不太留意。从 《Go 程序设计语言》中再一次体会到基本数据类型的别致，也对日后开发过程中，对小小的变量心存敬畏。

## 整型
对于 Go 中的 int 根据不同环境，甚至不同编译器实现的 int 类型表示都不一样。这样的处理想来可能是增加灵活性。  

对于确切的变量类型。比如，int8, uint8 这种类型表示是统一的。

需要注意的是无符号和有符号整型在循环中的使用。当超出类型表示范围无符号和有符号整型的溢出处理是不一样的。示例代码如下：
```
var b int8 = 127
fmt.Println(unsafe.Sizeof(b), b+1)

var c uint8
c = 255
fmt.Println(unsafe.Sizeof(c), c+1)
```

## 浮点型
浮点型只有两种 float32 和 float64。float64 精度更高，使用浮点型要注意精度问题。程序示例中从精度需求出发尽量使用 float64 类型。示例如下：
```
var e float64 = 16777216
fmt.Println(unsafe.Sizeof(e), e == e+1)     // true

var f byte
f = c
fmt.Println(unsafe.Sizeof(f), f)        // false
```

## 常量
常量比较有意思，分有类型常量和无类型常量。

有类型常量，比如 `const pi float32 = 3.14` 中 pi 即为有类型常量，有类型常量存储在只读内存区，数据段受保护，只读。

这里无类型常量是什么呢？  
3.14 即为无类型常量。上例中无类型常量 `3.14` 显式转换为 float32 类型的常量 pi（可以理解成 pi 是标记，通过 pi 可以查找到内存中的 3.14，3.14 存储在 32 位 bit 空间里）。

上例中是显式转换，如下代码：
```
var a float32
a = 2
```

其中，无类型常量 2 被隐式转换为 float32 类型的变量 a。因此，这里常量 2 不再是常量，也不在只读内存区中存储。

## 编译和运行时
```
var b int8 = 127
fmt.Println(unsafe.Sizeof(b), b+1)
// cannot use 65535 (untyped int constant) as int8 value in assignment (overflows)
// b = 65535
```

上述代码中，当为 b 赋值 128 是编译会报错，这类错误在编译阶段即会发现，类型溢出了。
而换种方式，将 b + 1 则能顺利通过编译，将程序执行溢出转到运行时处理。

## 字符串
字符串是不可更改的“串”，Go 的字符串通过结构体表示，该结构体由两个 16 Byte 组成。
更多信息可参考： [Go 设计与实现 - 字符串](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-string/)

# unicode
最后稍微过了下 unicode 和 utf-8。要注意的是 i 的 utf-8 字节表示，以及 utf-8 编码。
```
var i = "严"
fmt.Println(unsafe.Sizeof(i), i)
```

更多信息可参考: [字符编码笔记：ASCII，Unicode 和 UTF-8](https://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)

