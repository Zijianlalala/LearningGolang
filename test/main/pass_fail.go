package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)
func getFloat() (float64, error) {
		
}

func main(){
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, _ := strconv.ParseFloat(input, 64)
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is",status)
}