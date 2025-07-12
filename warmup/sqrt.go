package main

import "fmt"

type Solution struct{}

func (s Solution) mySqrt(x int) int {
	// TODO: Write your code here
	l, r := 0, x
	for l <= r {
		mid := (l + r) / 2
		midSqr := mid * mid
		if midSqr > x {
			r = mid - 1
		} else if midSqr < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r
}

func main() {
	solution := Solution{}

	input1 := 4
	expectedOutput1 := 2
	result1 := solution.mySqrt(input1)
	fmt.Println(result1 == expectedOutput1) // Expected output: true

	input2 := 8
	expectedOutput2 := 2
	result2 := solution.mySqrt(input2)
	fmt.Println(result2 == expectedOutput2) // Expected output: true

	input4 := 2
	expectedOutput4 := 1
	result4 := solution.mySqrt(input4)
	fmt.Println(result4 == expectedOutput4) // Expected output: true

	input5 := 3
	expectedOutput5 := 1
	result5 := solution.mySqrt(input5)
	fmt.Println(result5 == expectedOutput5) // Expected output: true

	input6 := 15
	expectedOutput6 := 3
	result6 := solution.mySqrt(input6)
	fmt.Println(result6 == expectedOutput6) // Expected output: true
}
