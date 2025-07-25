package slidingwindow

import (
	"fmt"
	"math"
)

type Solution struct{}

func (s *Solution) findSubstring(str string, pattern string) string {
	// TODO: Write your code here
	var strFreq [26]int
	var patternFreq [26]int
	for _, char := range pattern {
		patternFreq[char-'a']++
	}

	startPos := 0
	smallestWindow := ""
	smallestWindowSize := math.MaxInt32
	for endPos := 0; endPos < len(str); endPos++ {
		strFreq[str[endPos]-'a']++
		for contains(strFreq, patternFreq) {
			currWindowSize := (endPos - startPos + 1)
			if smallestWindowSize > currWindowSize {
				smallestWindowSize = currWindowSize
				smallestWindow = str[startPos : endPos+1]
			}
			strFreq[str[startPos]-'a']--
			startPos++
		}
	}

	return smallestWindow
}

func contains(strFreq, patternFreq [26]int) bool {
	for i := range strFreq {
		if patternFreq[i] > 0 && strFreq[i] < patternFreq[i] {
			return false
		}
	}
	return true
}

func main() {
	sol := Solution{}
	fmt.Println(sol.findSubstring("aabdec", "abc"))
	fmt.Println(sol.findSubstring("aabdec", "abac"))
	fmt.Println(sol.findSubstring("abdbca", "abc"))
	fmt.Println(sol.findSubstring("adcad", "abc"))
}
