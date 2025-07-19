package twopointer

import (
	"fmt"
	"sort"
)

type Solution struct{}

func twoPointer(arr []int, target, i int, triplets *[][]int) (int, int) {
	l, r := i+1, len(arr)-1
	for l < r {
		if arr[l]+arr[r] == target {
			*triplets = append(*triplets, []int{arr[i], arr[l], arr[r]})
			l++
			r--
			for l < r && arr[l] == arr[l-1] {
				l++
			}
			for l < r && arr[r] == arr[r+1] {
				r--
			}
		} else if arr[l]+arr[r] < target {
			l++
		} else {
			r--
		}
	}
	return -1, -1

}

func (s *Solution) searchTriplets(arr []int) [][]int {
	triplets := make([][]int, 0)
	// TODO: Write your code here
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		twoPointer(arr, -(arr[i]), i, &triplets)
	}
	return triplets
}

func main() {
	sol := Solution{}
	fmt.Println(sol.searchTriplets([]int{-3, 0, 1, 2, -1, 1, -2}))
	fmt.Println(sol.searchTriplets([]int{-5, 2, -1, -2, 3}))
}
