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

### Basic

map本身是一个无序的键值对集合

```go
// 创建map
// 声明map，但不会自动创建map
var myMap map[string][int]
// 通过make函数创建map
myMap = make(map[string]int)
// 简写
ranks := make(map[string]int)
// map字面量
myMap := map[string]int{"a":1 ,"b":2}
// 创建空map
myMap := map[string]int{}

// 操作
// 添加和访问键值对
ranks["gold"] = 1
// 如果访问一个未知key对应的value，value是零值，所以添加key-value对时，不需要其他额外操作（类似java中的put）
val := ranks["unknownKey"]
// 问题来了，如何判断某key对应的value是默认零值还是已赋零值，目的是判断key存不存在
// 访问value时返回多个值！判断ok的布尔值
val, ok = ranks["key"]
// 删除键值对
delete(ranks, "key")

// 遍历map
// 可以直接打印map
fmt.Println(ranks)
// 使用for...range loops 注意遍历顺序是无序的
for key, value := range ranks {
  // code
}
// 只用key值
for key := range ranks {
  // code
}
// 只用value值
for _, val := range ranks {
  // code
}
// 想要有序遍历map，首先要把key值取出来存到slice中，然后对slice排序，在根据slice的顺序从map中取值，可以实现按照某种顺序遍历map

```



## Chapter 8 Structs

C语言中的结构体？

### basic

```go
// 声明时就会赋零值给到变量myStruct中
var myStruct struct {
  name string
  number int
} 

// 通过点运算符访问结构体
// 既可以赋值又可以访问
myStruct.number = 1
fmt.Println(myStruct.number)

// 可以通过结构体字面量来创建结构体变量
myCar := car{name:"xxx", topSpeed: 337}

```

### 自定义类型

```go
// 声明格式
type myType struct {
  // field here
}
```

类型定义内部可以定义函数，但函数只能在当前作用域下使用，不能在定义外被调用。所以通常会将有关函数定义在类型定义外部，在package作用域下。



点运算符作用在结构体指针类型上时，与直接作用在结构体类型上时一样使用。



因为go的函数参数是值传递，所以当一个结构体过大时，通过值传递（即复制一个结构体）会占用额外资源，所以通常使用指针来作为函数的参数



结构体想要暴露给其他包使用，首字母要大写，结构体中的字段名称也要大写



### 匿名字段 anonymous fields

定义结构体时，可以使用匿名字段

```
type A struct {
	number int
}
type B struct {
	name string
	A // anonymous filed
}
func main() {
	b := B {name: "bbb"}
	b.number = 1 // 可以直接访问number字段
}
```

## Chapter 9 Defined type

上一章是基于结构体的自定义类型，本章讨论的自定义类型是基于基本类型的

基于基本类型的自定义类型**目的是**使变量分类更清晰，从语言结构上将两种变量分开，从语言特性上保证安全性

`method`是一种与某特定类型的值有关的`function`

### 转换

1. 基本类型转换为自定义类型
2. 自定类型1转换为自定义类型2（两个基于同一种基本类型）

### 操作符

自定义类型支持和他基于的类型所有相同的操作

### Method

目的是直接用`value`调用方法，方便？应用：数据类型转换

与`function`只有一个区别：你需要添加一个额外的参数——一个**接收参数**，在函数名前面

```go
// 例子
type MyType string

func main()  {
	value := MyType("a MyType value")
	value.sayHi()
}
// method
func (m MyType) sayHi()  {
  // 接收参数与普通参数几乎没有区别，都可以在函数体内被访问
	fmt.Println("Hi",m)
}
```

接收参数的命名规范

1. 一般参数名使用自定义类型的首字母小写



`method`几乎和`function`一样，可以接收参数，可以有返回值



#### 指针类型的接受者参数

将普通变量参数改为指针类型参数即可在函数内修改之



函数不能重载，但是方法可以重载（不同的自定义类型）



## Chapter 10 封装和嵌入

### Setter methods

结构体类型变量也是自定义类型变量，也有Method方法，且接受着参数需要是指针类型。

初始化结构体类型变量时，直接赋值无法保证初始化的值是否合法，可以通过Setter Method赋值来检查类型是否合法。

###  Getter methods

命名规范：直接使用变量名即可，首字母大写。



### Method Promotion

> 如果`struct A`属`struct B`，那么可以直接通过`B`的`value`访问`A`的`method`



## Chapter 11 Interfaces

### basic

```go
type myInterface interface {
  methodWithoutParameters()
  methodWithParameter(float64)
  methodWithReturnValue() string
}
// 一个type要satisfy interface
type MyType int

func (m MyType) methodWithoutParameters() {
  // code
}
// other interface methods satisfy
```

1. 在`Go`中不叫`implements`接口，而是`satisfy`接口

2. 一个类型既然满足了接口就要满足接口中的所有方法
3. 一个类型可以满足多个接口，一个接口可以被多个类型满足

```go
// 应用
var value myInterface // 声明接口类型变量
value = MyType // 将具体类型赋值给value
value.methodWithoutParameters()

```

### 具体类型&接口类型

具体类型：

1. method
2. value

接口类型：

1. method

### 类型断言 type assertions

存在的问题：接口类型变量只能调用接口声明的方法，不能调用具体类型中接口未声明的方法

`Type assertion`类似于`type conversion`

```go
var value MyInterface = MyType("xxx")
// type assertion
var m MyType = value.(MyType)
```

#### 转换失败

根据返回值判断是否转换成功

```go
myType, ok := value.(MyType)
```

### error

`error`也是接口类型变量

### Empty interface

```go
type Anything interface { 
}
```

 用于接收任何类型的值，有点类似于`Java`中的`Object`类？

## Chapter 12 错误处理

### `defer`关键字：

确保一个函数调用的发生，即便它很早退出（只能用于函数或者方法）

`defer`会在`panic`前调用



### `panic`

`panic`适用于处理bug/exception的，而不是用户的`mistake`

调用`panic`仍然会使程序崩溃，所以不是处理错误的理想方法

### `recover`函数

可以停止程序从`panic`中

用法：将`recover`放到一个函数中，然后`defer`该函数

结果

1. 当前函数中`panic`后的代码不会执行
2. `panic`所在的函数返回后，后面的函数正常执行

如果没有`panic`时，调用`recover`会返回`nil`

但当`panic`发生时，recover会返回传入`panic`中的参数

