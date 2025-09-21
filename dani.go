package main

import "fmt"

func sumAndAverage(num int) (int, float64) {
	var count int
	var sum int

	for num > 0 {

		sum += num % 10

		count++
		num = num / 10

	}
	return sum, float64(sum) / float64(count)

}

func main() {
	var number int
	var sum int = 0
	var avg float64 = 0

	fmt.Print("Enter a num: ")
	fmt.Scan(&number)

	sum, avg = sumAndAverage(number)
	fmt.Println("the sum is ", sum, "the avg is", avg)

}
