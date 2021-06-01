package main

import (
	"fmt"
	//"time"
	"strings"
)

func main() {
	// var now time.Time = time.Now();
	// var year int = now.Year();
	// fmt.Println(year)
	broken := "G#AAA#"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
