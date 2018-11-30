package main

import (
	"fmt"
)

func main() {

	for x := 1; x <= 21; x++ {
		if x%15 == 0 {
			fmt.Print(" DivisibleByFifteen")
		} else if x%5 == 0 {
			fmt.Print(" DivisibleByFive")
		} else if x%3 == 0 {
			fmt.Print(" DivisibleByThree")
		} else {
			fmt.Print(" ", x)
		}
	}

	fmt.Println()

	for x := 1; x <= 21; x++ {
		fmt.Print(" ", x)
		if x%10 == 0 {
			fmt.Print(" AfterDivisibleByTen")
		} else if x%5 == 0 {
			fmt.Print(" AfterDivisibleByFive")
		}
	}
}
