package slidingwindow

import "fmt"

type Solution struct{}

func (s *Solution) findStringAnagrams(str string, pattern string) []int {
	resultIndices := make([]int, 0)
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
				resultIndices = append(resultIndices, startPos)
			}
			strFreq[str[startPos]-'a']--
			startPos++
		}
	}
	return resultIndices
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
	fmt.Println(sol.findStringAnagrams("ppqp", "pq"))
	fmt.Println(sol.findStringAnagrams("abbcabc", "abc"))
}
