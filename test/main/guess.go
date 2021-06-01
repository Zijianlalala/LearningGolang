// guess challengs players to guess a random number.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"time"
)

func main()  {
	// variables
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	num := rand.Intn(100)+1
	reader := bufio.NewReader(os.Stdin)
	success := false

	fmt.Println("I've choose a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	for x:=10; x>0; x--{
		// read a number
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		grade, err := strconv.Atoi(input)
		if err != nil{
			log.Fatal(err)
		}
		// comparsion
		if grade > num {
			fmt.Println("it's bigger than random num! Only",x-1,"times left!")
		} else if grade < num {
			fmt.Println("it's smaller than random num!",x-1,"times left!")
		} else {
			success = true;
			fmt.Println("Congrats! You win!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was",num)
	}
}