package twopointer

import (
	"fmt"
	"math"
	"sort"
)

// Solution struct to encapsulate the functionality
type Solution struct{}

// searchTriplet method converted from Java to Go
func (s *Solution) searchTriplet(arr []int, targetSum int) int {
	// TODO: Write your code here
	sort.Ints(arr)
	closestSum := math.MaxInt32
	smallestDiff := math.MaxInt32
	for i := 0; i < len(arr)-2; i++ {
		twoPointer(arr, targetSum, arr[i], i+1, &closestSum, &smallestDiff)
	}
	return closestSum
}

func twoPointer(arr []int, target, currEle, l int, closestSum, smallestDiff *int) {
	r := len(arr) - 1
	for l < r {
		currSum := arr[l] + arr[r] + currEle
		currDiff := abs(currSum - target)

		if currSum == target {
			*closestSum = currSum
			*smallestDiff = 0
			return
		}

		if currDiff < *smallestDiff {
			*smallestDiff = currDiff
			*closestSum = currSum
		} else if currDiff == *smallestDiff {
			if currSum < *closestSum {
				*closestSum = currSum
			}
		}

		if currSum < target {
			l++
		} else if currSum > target {
			r--
		}
	}
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
	fmt.Println(solution.searchTriplet([]int{-1, 0, 2, 3}, 2))
	fmt.Println(solution.searchTriplet([]int{-3, -1, 1, 2}, 1))
	fmt.Println(solution.searchTriplet([]int{1, 0, 1, 1}, 100))
	fmt.Println(solution.searchTriplet([]int{0, 0, 1, 1, 2, 6}, 5))
}
