package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Define Solution struct
type Solution struct{}

// Function to check if given sentence is pangram
func (s *Solution) checkIfPangram(sentence string) bool {
	// TODO: Write your code here
	charMap := make(map[rune]struct{})
	sentence = strings.ToLower(sentence)

	for _, char := range sentence {
		_, ok := charMap[char]
		if unicode.IsLetter(char) && !ok {
			charMap[char] = struct{}{}
		}
	}
	return len(charMap) == 26
}

func main() {
	sol := Solution{}

	// Test cases
	fmt.Println(sol.checkIfPangram("TheQuickBrownFoxJumpsOverTheLazyDog"))                  // Expected output: true
	fmt.Println(sol.checkIfPangram("This is not a pangram"))                                // Expected output: false
	fmt.Println(sol.checkIfPangram("abcdef ghijkl mnopqr stuvwxyz"))                        // Expected output: true
	fmt.Println(sol.checkIfPangram(""))                                                     // Expected output: false
	fmt.Println(sol.checkIfPangram("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")) // Expected output: true
}
