package twopointer

import "fmt"

// Solution struct to mimic a class.
type Solution struct{}

func (s *Solution) makeSquares(arr []int) []int {
	n := len(arr)
	squares := make([]int, n)

	l, r := 0, len(arr)-1
	resIdx := n - 1
	for l <= r {
		lSqr := arr[l] * arr[l]
		rSqr := arr[r] * arr[r]

		if lSqr > rSqr {
			squares[resIdx] = lSqr
			l++
		} else {
			squares[resIdx] = rSqr
			r--
		}
		resIdx--
	}

	return squares
}

func main() {
	sol := Solution{}
	result := sol.makeSquares([]int{-2, -1, 0, 2, 3})
	for _, num := range result {
		fmt.Print(num, " ")
	}
	fmt.Println()

	result = sol.makeSquares([]int{-3, -1, 0, 1, 2})
	for _, num := range result {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
