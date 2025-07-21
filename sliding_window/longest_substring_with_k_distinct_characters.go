package slidingwindow

import "fmt"

type Solution struct{}

func (s *Solution) findLength(str string, k int) int {
	// TODO: Write your code here
	startPos := 0

	largestLen := 0
	charMap := make(map[byte]int)
	for endPos := 0; endPos < len(str); endPos++ {
		currChar := str[endPos]
		if freq, has := charMap[currChar]; has {
			charMap[currChar] = freq + 1
		} else {
			charMap[currChar] = 1
		}

		for len(charMap) > k {
			currChar = str[startPos]
			charMap[currChar] = charMap[currChar] - 1

			if count, _ := charMap[currChar]; count == 0 {
				delete(charMap, currChar)
			}
			startPos++
		}
		largestLen = max(largestLen, endPos-startPos+1)

	}
	return largestLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sol := Solution{}
	fmt.Println("Length of the longest substring:", sol.findLength("araaci", 2))
	fmt.Println("Length of the longest substring:", sol.findLength("araaci", 1))
	fmt.Println("Length of the longest substring:", sol.findLength("cbbebi", 3))
}
