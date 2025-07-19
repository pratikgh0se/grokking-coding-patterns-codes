package twopointer

import (
	"fmt"
	"sort"
)

// Solution struct to simulate a class
type Solution struct{}

// searchTriplets method as per the Java version
func (s *Solution) searchTriplets(arr []int, target int) int {
	count := 0
	sort.Ints(arr)
	// TODO: Write your code here
	for i := 0; i < len(arr)-2; i++ {
		twoPointer(arr, arr[i], i+1, target, &count)
	}
	return count
}

func twoPointer(arr []int, currEle, l, target int, count *int) {
	r := len(arr) - 1
	for l < r {
		currSum := arr[l] + arr[r] + currEle
		if currSum < target {
			*count += (r - l)
			l++
		} else {
			r--
		}
	}
}

func main() {
	sol := Solution{}
	fmt.Println(sol.searchTriplets([]int{-1, 0, 2, 3}, 3))
	fmt.Println(sol.searchTriplets([]int{-1, 4, 2, 1, 3}, 5))
}
