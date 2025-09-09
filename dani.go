package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {

		for z := 1; z <= 10; z++ {

			fmt.Println(i, "X", z, "= ", z*i)

		}

	}

}
