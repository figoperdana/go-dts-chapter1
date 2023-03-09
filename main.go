package main

import (
	"fmt"
)

func main() {
	word := "selamat malam"
	chars := (word)
	charCount := make(map[string]int)

	for _, char := range word {
		charCount[string(char)]++
	}

	for i := 0; i < len(chars); 
	i++ {
		char := string(chars[i])
		println(char)
	}
	fmt.Println()
	fmt.Println(charCount)
}