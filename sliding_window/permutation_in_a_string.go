package slidingwindow

import "fmt"

type Solution struct{}

func (s *Solution) findPermutation(str string, pattern string) bool {
	// TODO: Write your code here
	var strFreq [26]int
	var patternFreq [26]int
	k := len(pattern)
	for _, char := range pattern {
		patternFreq[char-'a']++
	}

	startPos := 0
	for endPos := 0; endPos < len(str); endPos++ {
		strFreq[str[endPos]-'a']++
		if startPos+k-1 == endPos {
			if isEqual(strFreq, patternFreq) {
				return true
			}
			strFreq[str[startPos]-'a']--
			startPos++
		}
	}

	return false
}

func isEqual(strFreq, patternFreq [26]int) bool {
	for i := range strFreq {
		if strFreq[i] != patternFreq[i] {
			return false
		}
	}
	return true
}

func main() {
	sol := Solution{}
	fmt.Println("Permutation exist:", sol.findPermutation("oidbcaf", "abc"))
	fmt.Println("Permutation exist:", sol.findPermutation("odicf", "dc"))
	fmt.Println("Permutation exist:", sol.findPermutation("bcdxabcdy", "bcdyabcdx"))
	fmt.Println("Permutation exist:", sol.findPermutation("aaacb", "abc"))
}
