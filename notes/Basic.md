# Notes

## Chapter 1 Syntax Basics

1. 程序结构
   1. package
   2. import
   3. code
2. 变量类型
   1. int
   2. int32
   3. float64
   4. bool
   5. string
3. 声明变量
   1. var name type = value
   2. var name = value
   3. name := value
4. 变量命名规则
   1. 首字母大写：public
   2. 首字母小写：protected
   3. 驼峰命名规则
5. 类型转化 `type(value)`

## Chapter 2 Conditionals and Loops

1. 函数可以多个返回值，常用的返回多个值会包括错误状态码
2. go要求每个声明的变量必须被使用
   1. 如果不想使用，则用`'_'`替代变量名
   2. 处理返回值
3. `log.Fatal(err)`会直接终止程序
4. 条件语句，不需要将布尔表达式用圆括号括起来
5. 变量名如果与关键字重复，会覆盖掉关键字的语义
6. 功能包
   1. `strings.TrimSpace("xxx")`，去掉前后空格
   2. `strconv.ParseFloat(string, 64)`，将字符串转换为浮点数
7. 代码段（Blocks），变量只在当前代码段内有定义（作用域）
8. `:=`左边的变量如果已经存在则表示赋值，如果不存在，则表示声明
9. 循环`for x:=0; x < 10; x++ {}`

完成了一个`guess game`

## Chapter 3 Functions

1. 浮点数的格式化输出`fmt.printf("%0.2f", num)`，其中`%0.2f`里面`%f`表示浮点数，小数点前的数字表示整个数字的最小长度，小数点后的数字表示保留几位小数
2. 框架

```go
func name(param type) returnType {
  // code
  return value
}
```

3. 函数返回`error`类型表示执行状态
4. 函数的形参接收到的是实参的拷贝，也就是**值传递**
5. **指针**
   1. 声明指针类型变量 `*int`
   2. 获取变量的地址`&variable`
   3. 获取指针变量对应的值`*pointer`
   4. 更新指针变量对应的值`*pointer = 3`

## Chapter 4 Packages

### 项目结构

* `workspace`
  * `src`
  * `bin`
  * `pkg`

新建项目目录结构后，碰到了项目根目录找不到的问题，解决

```bash
go env -w GOPATH=/Users/i526563/go
go env -w GO111MODULE=off
```



### Package命名规则

1. 全部小写
2. 简写且含义明确
3. 尽可能一个单词，如果多个，首字母也不大写，例如`strconv`
4. 命名不要和常规单词冲突

### 常量 Constant

```go
const Days int = 7
```



1. 使用`const`关键字声明
2. 声明时必须赋值

### 设置环境变量 GOPATH

设置

```bash
export GOPATH="/code"
```



### 包的下载

```bash
go get github.com/headfirstgo/greeting
```

会将远程包自动放到`workspace`中，`import`之后就可以使用了

阅读包相关文档命令

```bash
go doc packagePath
```

另外在源码中，首行的注释和函数前的注释就可被go doc读取，且注释也有规范

1. 需要是完整的句子
2. 包注释以`Package`开头
3. 函数注释以`Function Name`开头

## Chapter 5 Arrays

### Basic

声明数组

```go
var myArray [4]string
```

* 初始值是零值
* 数组字面量，`nums :=[3]int{1,2,4}`
* 遍历时，防止数组越界，可以用`len(array)`，也可以用`for...range`

```go
nums := [4]int{1,2,3,4}	
// 如果不需要index，可以使用'_'替代index
for index, value := range nums {
	fmt.Println(index,value)
}

```

### 读取本地文件

```go
package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
```

## Chapter 6 Slices

> 为了解决数组固定大小的问题

`slice`是一个可以动态增长的集合类型，使用`array`实现的。就像`Java`中的`ArrayList`底层也是由`Array`实现的

```go
// 不指定元素个数时，创建的是slice类型的变量
var notes []string // 声明slice
notes = make([]string, 7) //创建slice
// 或者简写成
notes := make([]string, 7)
// 或者使用字面量
notes := []int{1,2,3}
// 通过操作数组创建slice (from 1 to 3(exclusive))
mySlice := myArray[1:3]
```



### Append方法

```go
mySlice = append(mySlice, item)
```

### 命令行参数

`os.Args`返回的是`slice`类型

### variadic function

可以接收可变数量的参数

```go
func severalInts(numbers ...int) {
	// code
}
// 但是不能直接把slice类型变量传到该类函数中，需要使用特殊符号进行传递
func main() {
  initSlice := []int{1,2,3}
  severalInts(initSlice...)
}
```

## Chapter 7 Maps

