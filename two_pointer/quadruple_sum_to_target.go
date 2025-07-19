package twopointer

import (
	"fmt"
	"sort"
)

type Solution struct{}

func (s *Solution) searchQuadruplets(arr []int, target int) [][]int {
	sort.Ints(arr)
	var quadruplets [][]int
	n := len(arr) - 1
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			twoPointer(arr, j+1, n, arr[i], arr[j], target, &quadruplets)
		}
	}
	// TODO: Write your code here
	return quadruplets
}

func twoPointer(arr []int, l, r, q1, q2, target int, res *[][]int) {
	for l < r {
		currSum := q1 + q2 + arr[l] + arr[r]
		if currSum == target {
			*res = append(*res, []int{q1, q2, arr[l], arr[r]})
			l++
			r--
			for l < r && arr[l] == arr[l-1] {
				l++
			}
			for l < r && arr[r] == arr[r+1] {
				r--
			}
		} else if currSum < target {
			l++
		} else {
			r--
		}
	}
}

func main() {
	sol := Solution{}
	fmt.Println(sol.searchQuadruplets([]int{4, 1, 2, -1, 1, -3}, 1))
	fmt.Println(sol.searchQuadruplets([]int{2, 0, -1, 1, -2, 2}, 2))
}
