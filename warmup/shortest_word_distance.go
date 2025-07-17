package warmup

import (
	"fmt"
)

type Solution struct{}

// shortestDistance finds the shortest distance between two words in an array.
func (s *Solution) shortestDistance(words []string, word1 string, word2 string) int {
	// TODO: Write your code here
	shortestDistance := 999999
	word1Idx, word2Idx := -999999, 999999
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			word1Idx = i
		}
		if words[i] == word2 {
			word2Idx = i
		}
		shortestDistance = min(shortestDistance, abs(word1Idx-word2Idx))
	}
	return shortestDistance
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	solution := Solution{}

	// Test case 1
	words1 := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word11 := "fox"
	word21 := "dog"
	expectedOutput1 := 5
	actualOutput1 := solution.shortestDistance(words1, word11, word21)
	fmt.Println("Test case 1:", expectedOutput1 == actualOutput1)

	// Test case 2
	words2 := []string{"a", "b", "c", "d", "a", "b"}
	word12 := "a"
	word22 := "b"
	expectedOutput2 := 1
	actualOutput2 := solution.shortestDistance(words2, word12, word22)
	fmt.Println("Test case 2:", expectedOutput2 == actualOutput2)

	// Test case 3
	words3 := []string{"a", "c", "d", "b", "a"}
	word13 := "a"
	word23 := "b"
	expectedOutput3 := 1
	actualOutput3 := solution.shortestDistance(words3, word13, word23)
	fmt.Println("Test case 3:", expectedOutput3 == actualOutput3)

	// Test case 4
	words4 := []string{"a", "b", "c", "d", "e"}
	word14 := "a"
	word24 := "e"
	expectedOutput4 := 4
	actualOutput4 := solution.shortestDistance(words4, word14, word24)
	fmt.Println("Test case 4:", expectedOutput4 == actualOutput4)
}
