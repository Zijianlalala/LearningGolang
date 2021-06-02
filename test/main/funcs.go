package main

import(
	"fmt"
	"errors"
)

func main()  {
	amount := 6;
	fmt.Println(amount)
	fmt.Println(&amount)
}

func sayHi(){
	fmt.Println("Hi~")
}

func repeatLines(line string, times int){
	for x:=0; x<times; x++{
		fmt.Println(line)
	}
}

func double(number float64) float64  {
	return number * 2
}

func paintNeeded(height float64, width float64) (float64, error)  {
	if height < 0 || width < 0 {
		return 0, errors.New("cant be negative.")
	}
	return height * width /10, nil
}