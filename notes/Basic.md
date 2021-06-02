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