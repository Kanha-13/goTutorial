package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")
	fmt.Println("What is your age?")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	fmt.Println("Ok so your are ", age, "years old \n")
	fmt.Println("This variable is of type %T \n", age)
}
