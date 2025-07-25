package slidingwindow

import "fmt"

type Solution struct{}

func (s Solution) findLength(arr []int, k int) int {
	// TODO: Write your code here
	startPos := 0
	longestSubArray := 0
	var freq [2]int
	for endPos := 0; endPos < len(arr); endPos++ {
		currInt := arr[endPos]
		freq[currInt]++
		if freq[0] > k {
			currInt := arr[startPos]
			startPos++
			freq[currInt]--
		}
		longestSubArray = max(longestSubArray, endPos-startPos+1)

	}
	return longestSubArray
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sol := Solution{}
	fmt.Println(sol.findLength([]int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}, 2))
	fmt.Println(sol.findLength([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 3))
}
