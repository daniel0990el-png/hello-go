package main

import "fmt"

func main() {
	var num int
	var x int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

	sum := 0
	x = 0

	for i := 1; num > 0; i++ {

		x = num % 10
		num = num / 10

		fmt.Println(x)
		sum += x
	}
	fmt.Print(sum)
}
