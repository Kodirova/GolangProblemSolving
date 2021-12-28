package main

import (
	"fmt"
)

func FizzBuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(fizz + buzz)
	} else if i%3 == 0 {
		fmt.Println(fizz)
	} else if i%5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(i)
	}

}

func main() {
	var input int
	fmt.Println("Enter number: ")
	fmt.Scanf("%d", &input)
	for i := 1; i < input; i++ {
		FizzBuzz(i)
	}

}
