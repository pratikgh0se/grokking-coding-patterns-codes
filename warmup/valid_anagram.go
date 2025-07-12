package main

import (
	"fmt"
)

type Solution struct{}

// isAnagram checks if string t is an anagram of string s.
// It first compares the lengths of the strings.
// If they are equal, it uses a map to count the frequency of characters in both strings.
// It then checks if all characters have a frequency of 0, indicating that t is an anagram of s.
func createCharCountArray(s string) [128]int {
	var arr [128]int
	for _, char := range s {
		arr[char]++
	}
	return arr
}

func (sol *Solution) isAnagram(s, t string) bool {
	// TODO: Write your code here
	sCountArr := createCharCountArray(s)
	tCountArr := createCharCountArray(t)
	return sCountArr == tCountArr
}

func main() {
	sol := Solution{}

	// Test case 1
	s1 := "listen"
	t1 := "silent"
	fmt.Println(sol.isAnagram(s1, t1)) // Expected output: true

	// Test case 2
	s2 := "hello"
	t2 := "world"
	fmt.Println(sol.isAnagram(s2, t2)) // Expected output: false

	// Test case 3
	s3 := "anagram"
	t3 := "nagaram"
	fmt.Println(sol.isAnagram(s3, t3)) // Expected output: true

	// Test case 4
	s4 := "rat"
	t4 := "car"
	fmt.Println(sol.isAnagram(s4, t4)) // Expected output: false

	// Test case 5
	s5 := ""
	t5 := ""
	fmt.Println(sol.isAnagram(s5, t5)) // Expected output: true
}
