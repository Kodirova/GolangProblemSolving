package main

import (
	"fmt"
)

func fibonaccirecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonaccirecursion(n-1) + fibonaccirecursion(n-2)
}

func fibonnacislice(n int) int {
	sequence := make([]int, n+1, n+2)
	if n < 2 {
		sequence = sequence[0:2]
	}
	sequence[0] = 0
	sequence[1] = 1
	for i := 2; i <= n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
	return sequence[n]
}

func Fibonacciloop(n int) int {
	n1, n2 := 0, 1
	var n3 int
	for i := 2; i <= n; i++ {
		n3 = n1 + n2
		n1 = n2
		n2 = n3

	}
	return n3

}

func main() {
	r := fibonaccirecursion(25)
	s := fibonnacislice(25)
	l := Fibonacciloop(25)

	fmt.Printf("Fibonacci sequence using recursion is = %d \n", r)
	fmt.Printf("Fibonacci sequence using slice is = %d \n", s)
	fmt.Printf("Fibonacci sequence using loop is = %d \n", l)
}
