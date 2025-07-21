package slidingwindow

import (
	"fmt"
	"math"
)

type Solution struct{}

func (s *Solution) findMinSubArray(S int, arr []int) int {
	// TODO: Write your code here
	startPos := 0
	curSum := 0
	smallestLen := math.MaxInt32

	for endPos := 0; endPos < len(arr); endPos++ {
		curSum += arr[endPos]
		for curSum >= S {
			smallestLen = min(smallestLen, endPos-startPos+1)
			curSum = curSum - arr[startPos]
			startPos++
		}
	}

	if smallestLen == math.MaxInt32 {
		return 0
	}

	return smallestLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sol := Solution{}
	result := sol.findMinSubArray(7, []int{2, 1, 5, 2, 3, 2})
	fmt.Println("Smallest subarray length:", result)
	result = sol.findMinSubArray(7, []int{2, 1, 5, 2, 8})
	fmt.Println("Smallest subarray length:", result)
	result = sol.findMinSubArray(8, []int{3, 4, 1, 1, 6})
	fmt.Println("Smallest subarray length:", result)
}
