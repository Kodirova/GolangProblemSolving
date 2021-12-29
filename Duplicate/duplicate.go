package main

import "fmt"

func checkDuplicate(s []int) bool {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(s); i++ {
		if visited[s[i]] == true {
			return true
		} else {
			visited[s[i]] = true
		}
	}
	return false
}

func main() {
	fmt.Println(checkDuplicate([]int{1, 4, 7, 2, 2}))
}
