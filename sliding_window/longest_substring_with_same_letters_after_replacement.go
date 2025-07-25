package slidingwindow

import "fmt"

type Solution struct{}

func (s *Solution) findLength(str string, k int) int {
	// TODO: Write your code here
	startPos := 0
	longestSubArray := 0
	var charFrequency [26]int
	for endPos := 0; endPos < len(str); endPos++ {
		currChar := str[endPos]
		charFrequency[currChar-'a']++
		for ((endPos - startPos + 1) - findMax(charFrequency)) > k {
			currChar = str[startPos]
			startPos++
			charFrequency[currChar-'a']--
		}
		longestSubArray = max(longestSubArray, endPos-startPos+1)
	}
	return longestSubArray
}

func findMax(charFrequency [26]int) int {
	maxC := -1
	for _, ele := range charFrequency {
		maxC = max(maxC, ele)
	}
	return maxC
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sol := Solution{}
	fmt.Println(sol.findLength("aabccbb", 2))
	fmt.Println(sol.findLength("abbcb", 1))
	fmt.Println(sol.findLength("abccde", 1))
}
