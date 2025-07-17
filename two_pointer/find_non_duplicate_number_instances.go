package twopointer

import "fmt"

// Solution is a struct that doesn't have any fields, but it contains a method called "moveElements".
type Solution struct {
}

func (sol *Solution) moveElements(arr []int) int {
	nextNonDuplicate := 1

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx] != arr[idx-1] {
			nextNonDuplicate++
		}
	}
	// TODO: Write your code here
	return nextNonDuplicate
}

func main() {
	sol := Solution{}
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(sol.moveElements(arr))

	arr = []int{2, 2, 2, 11}
	fmt.Println(sol.moveElements(arr))
}
