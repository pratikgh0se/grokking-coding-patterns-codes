package main

import (
	"fmt"
)

type Solution struct{}

// numGoodPairs calculates the number of good pairs in the array.
// A pair (i, j) is good if nums[i] == nums[j] and i < j.
func (sol *Solution) numGoodPairs(nums []int) int {
	pairCount := 0
	// TODO: Write your code here
	pairMap := make(map[int]int)
	for _, ele := range nums {
		if apCount, has := pairMap[ele]; has {
			pairCount += apCount
		}
		pairMap[ele]++
	}
	return pairCount
}

func main() {
	sol := Solution{}

	nums1 := []int{1, 2, 3, 1, 1, 3}
	result1 := sol.numGoodPairs(nums1)
	fmt.Printf("Result 1: %d (Expected: 4)\n", result1)

	nums2 := []int{1, 1, 1, 1}
	result2 := sol.numGoodPairs(nums2)
	fmt.Printf("Result 2: %d (Expected: 6)\n", result2)

	nums3 := []int{1, 2, 3}
	result3 := sol.numGoodPairs(nums3)
	fmt.Printf("Result 3: %d (Expected: 0)\n", result3)
}
