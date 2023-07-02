package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter your age")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	fmt.Println("So you are ", age)

	converted_age, err := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("you will be", converted_age+1, "next year!")
	}

}
