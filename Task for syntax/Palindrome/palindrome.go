package main

import (
	"fmt"
)

func Palindrome(s string) bool {
	reversed := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	if s == string(reversed) {
		fmt.Printf("The word: %s is palindrome \n", s)
		return true
	} else {
		fmt.Printf("The word: %s is not palindrome \n", s)
		return false
	}

}

func checkpolindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			fmt.Printf("The word: %s is not palindrome \n", str)

			return false
		}
	}
	fmt.Printf("The word: %s is  palindrome \n", str)
	return true
}

func main() {
	var word string
	fmt.Println("Enter word to check:")
	fmt.Scanf("%s", &word)
	fmt.Println("1st attempt")
	Palindrome(word)
	fmt.Println("2nd attempt")
	checkpolindrome(word)
}
