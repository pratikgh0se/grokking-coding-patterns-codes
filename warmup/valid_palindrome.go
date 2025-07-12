package main

import (
	"fmt"
	"strings"
	"unicode"
)

// isLetterOrDigit checks if the given rune is a letter or a digit.
func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// isPalindrome checks if the given string is a palindrome,
// considering only letters and digits and ignoring cases.
func (sol *Solution) isPalindrome(s string) bool {
	r, l := 0, len(s)-1
	for r < l {
		for r < l && !isLetterOrDigit(rune(s[r])) {
			r++
		}
		for r < l && !isLetterOrDigit(rune(s[l])) {
			l--
		}
		if strings.ToLower(string(s[r])) != strings.ToLower(string(s[l])) {
			return false
		}
		r++
		l--
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(isPalindrome("A man, a plan, a canal, Panama!")) // true
	fmt.Println(isPalindrome("race a car"))                      // false
	fmt.Println(isPalindrome("Was it a car or a cat I saw?"))    // true
	fmt.Println(isPalindrome("Madam, in Eden, I'm Adam."))       // true
	fmt.Println(isPalindrome(""))                                // true
}
