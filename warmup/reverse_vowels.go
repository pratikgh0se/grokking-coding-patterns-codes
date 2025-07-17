package warmup

import (
	"fmt"
)

// Solution class equivalent in Go
type Solution struct{}

func swapRunes(str string, i, j int) string {
	r := []rune(str)
	r[i], r[j] = r[j], r[i]
	return string(r)
}

// reverseVowels function to reverse vowels in a string
func (s *Solution) reverseVowels(str string) string {
	// TODO: Write your code here
	vowels := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	right, left := 0, len(str)-1
	for right <= left {
		for right < left {
			if _, ok := vowels[str[right]]; ok {
				break
			}
			right++
		}
		for right < left {
			if _, ok := vowels[str[left]]; ok {
				break
			}
			left--
		}
		_, rok := vowels[str[right]]
		_, lok := vowels[str[left]]
		if rok && lok {
			str = swapRunes(str, right, left)
		}
		right++
		left--
	}
	return str
}

func main() {
	solution := Solution{}

	// Test Case 1
	s1 := "hello"
	expectedOutput1 := "holle"
	actualOutput1 := solution.reverseVowels(s1)
	fmt.Println("Test Case 1:", expectedOutput1 == actualOutput1)

	// Test Case 2
	s2 := "DesignGUrus"
	expectedOutput2 := "DusUgnGires"
	actualOutput2 := solution.reverseVowels(s2)
	fmt.Println("Test Case 2:", expectedOutput2 == actualOutput2)

	// Test Case 3
	s3 := "AEIOU"
	expectedOutput3 := "UOIEA"
	actualOutput3 := solution.reverseVowels(s3)
	fmt.Println("Test Case 3:", expectedOutput3 == actualOutput3)

	// Test Case 4
	s4 := "aA"
	expectedOutput4 := "Aa"
	actualOutput4 := solution.reverseVowels(s4)
	fmt.Println("Test Case 4:", expectedOutput4 == actualOutput4)

	// Test Case 5
	s5 := ""
	expectedOutput5 := ""
	actualOutput5 := solution.reverseVowels(s5)
	fmt.Println("Test Case 5:", expectedOutput5 == actualOutput5)
}
