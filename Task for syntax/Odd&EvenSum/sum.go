package main

import "fmt"

func Sum(n int) {
	odd_sum := 0
	even_sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even_sum += i
		} else {
			odd_sum += i
		}
	}
	fmt.Printf("Sum of even numbers = %d \n", even_sum)
	fmt.Printf("Sum of even numbers = %d \n", odd_sum)
}

func main() {
	var limit int
	fmt.Println("Enter number:")
	fmt.Scanf("%d", &limit)
	Sum(limit)

}
