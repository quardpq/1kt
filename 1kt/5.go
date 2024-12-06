package main

import (
	"fmt"
	"strings"
)

func areCharactersUnique(s string) bool {
	s = strings.ToLower(s)
	charMap := make(map[rune]bool)
	for _, char := range s {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}

func main() {
	fmt.Println(areCharactersUnique("abcd"))      // true
	fmt.Println(areCharactersUnique("abCdefAaf")) // false
	fmt.Println(areCharactersUnique("aabcd"))     // false
}
