package twopointer

import "fmt"

// Solution struct to hold methods
type Solution struct{}

// sort is the method to sort the array
// all elements < low are 0 and all elements > high are 2
// all elements from >= low < i are 1
func (s *Solution) sort(arr []int) []int {
	// TODO: Write your code here
	low := 0
	high := len(arr) - 1
	curr := 0
	for curr <= high {
		if arr[curr] == 0 {
			arr[curr], arr[low] = arr[low], arr[curr]
			low++
			curr++
		} else if arr[curr] == 2 {
			arr[curr], arr[high] = arr[high], arr[curr]
			high--
		} else if arr[curr] == 1 {
			curr++
		}

	}
	return arr
}

func main() {
	sol := Solution{}
	arr := []int{1, 0, 2, 1, 0}
	arr = sol.sort(arr)
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()

	arr = []int{2, 2, 0, 1, 2, 0}
	arr = sol.sort(arr)
	for _, num := range arr {
		fmt.Print(num, " ")
	}
}
