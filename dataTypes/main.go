package main

import "fmt"

var public = "token" //starting letter small means private var
var Public int = 9   //starting letter Capital means public var
const aconstantvar = 100

func main() {
	var username string = "kanha"
	fmt.Println(username)
	fmt.Printf("This is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("This is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("This is of type: %T \n", smallVal)

	var smallInt int = 256
	fmt.Println(smallInt)
	fmt.Printf("This is of type: %T \n", smallInt)

	var smallFloat float32 = 255.465465454654564
	fmt.Println(smallFloat)
	fmt.Printf("This is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.465465454654564
	fmt.Println(bigFloat)
	fmt.Printf("This is of type: %T \n", bigFloat)

	var value int
	fmt.Println(value)
	fmt.Printf("This is of type: %T \n", value)

	var value1 float32
	fmt.Println(value1)
	fmt.Printf("This is of type: %T \n", value1)

	var value2 string
	fmt.Println(value2)
	fmt.Printf("This is of type: %T \n", value2)

	var withouttype = "hello"
	fmt.Println(withouttype)
	fmt.Printf("This is of type: %T \n", withouttype)

	var withouttype1 = 1
	fmt.Println(withouttype1)
	fmt.Printf("This is of type: %T \n", withouttype1)

	novartype := "hello no var"
	fmt.Println(novartype)
	fmt.Printf("This is of type: %T \n", novartype)

}
