# goWelz
​	go语言自己的程序库，里面有自己写的代码，自己写的小工具，和一些方便自己使用的库。

### 代码介绍

#### 1. GoAinB

​	这是一个方便判断数组或字符串，是否包含某个元素的库。可以很好的帮助我去做判断。更好的实现代码复用。该库下面拥有3个方法。

1. 判断数组是否包含某个值 AinBArrayOneItem(item string, items []string) bool
2. 判断数组是否包含另一个数组，AinBArrayItems(itemA []string, itemB []string) bool
3. 判断字符串A是否被字符串B包含AinBStr(strA string, strB string) bool

使用方法如下：

```go
	strA := "Go is Good"
	strB := "Go"
	fmt.Println(goAinB.AinBStr(strB,strA))
```

#### 2. GoRoutineTest

​	这是一个Go语言的小实验，实验了一下使用Go Routine 结合生产者消费者模型的多线程的使用，使用Sync WaitGroup来进行Go Routine协程等待，确定Go Routine协程是否执行完毕。

​	第二版程序中，用锁实现了计数器，线程间可以通过获取变量的值来确定是否应该结束