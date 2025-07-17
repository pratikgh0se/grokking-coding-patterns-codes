package twopointer

import (
	"fmt"
)

type Solution struct {
}

// search - finds a pair in arr with a given targetSum.
func (sol *Solution) search(arr []int, targetSum int) []int {
	// TODO: Write your code here
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l]+arr[r] == targetSum {
			return []int{l, r}
		} else if arr[l]+arr[r] < targetSum {
			l++
		} else {
			r--
		}
	}
	return []int{-1, -1} // pair not found
}

func main() {
	result := search([]int{1, 2, 3, 4, 6}, 6)
	fmt.Printf("Pair with target sum: [%d, %d]\n", result[0], result[1])
	result = search([]int{2, 5, 9, 11}, 11)
	fmt.Printf("Pair with target sum: [%d, %d]\n", result[0], result[1])
}
