package main

import "fmt"

// maxDigit מחזירה את הספרה הכי גדולה במספר
func maxDigit(num int) int {
	if num < 0 {
		num = -num // אם המספר שלילי נהפוך אותו לחיובי
	}

	max := 0
	for num > 0 {
		digit := num % 10
		if digit > max {
			max = digit
		}
		num /= 10
	}
	return max
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	biggest := maxDigit(number)
	fmt.Println("The biggest digit is:", biggest)
}
