package main

import "fmt"

// Solution struct type
type Solution struct{}

// containsDuplicate checks for duplicates in a slice of integers
func (s *Solution) containsDuplicate(nums []int) bool {
	// ToDo: Write Your Code Here.
	seen := make(map[int]struct{})

	for _, val := range nums {
		if _, ok := seen[val]; ok {
			return true
		}
		seen[val] = struct{}{}
	}
	return false
}

func main() {
	solution := Solution{}

	nums1 := []int{1, 2, 3, 4}
	fmt.Println(solution.containsDuplicate(nums1)) // Expected output: false

	nums2 := []int{1, 2, 3, 1}
	fmt.Println(solution.containsDuplicate(nums2)) // Expected output: true

	nums3 := []int{}
	fmt.Println(solution.containsDuplicate(nums3)) // Expected output: false

	nums4 := []int{3, 2, 6, -1, 2, 1}
	fmt.Println(solution.containsDuplicate(nums4)) // Expected output: true
}
