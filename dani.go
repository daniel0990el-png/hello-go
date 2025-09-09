package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var age int
	age = 12

	fmt.Print("type your age: ")

	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("GOOD")
	} else {

		fmt.Println("BAD")
	}

}
