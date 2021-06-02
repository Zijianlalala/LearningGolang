package main

import(
	"fmt"
)

func main()  {
	num := 10
	changeNum(&num)
	fmt.Println(num)
}

func changeNum(num *int)  {
	*num *= 2
}