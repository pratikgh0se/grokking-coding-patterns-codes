// Solution struct
type Solution struct{}

// findMaxSumSubArray method to find the maximum sum of a subarray of size k
func (s *Solution) findMaxSumSubArray(k int, arr []int) int {
	maxSum := 0

	currSum := 0
	startPos := 0
	for endPos := 0; endPos < len(arr); endPos++ {
		currSum += arr[endPos]

		if (endPos - startPos) == k-1 {
			maxSum = max(maxSum, currSum)
			currSum = currSum - arr[startPos]
			startPos++
		}

	}
	// TODO: Write your code here
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := Solution{}
	fmt.Println("Maximum sum of a subarray of size K: ",
		s.findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println("Maximum sum of a subarray of size K: ",
		s.findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))
}