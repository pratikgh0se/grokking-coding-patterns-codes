package fastandslowpointer

import "fmt"

type Solution struct{}

func (s *Solution) find(num int) bool {
	// TODO: Write your code here
	slow, fast := num, num
	for {
		slow = getSquaredSum(slow)
		fast = getSquaredSum(getSquaredSum(fast))
		if slow == 1 || fast == 1 {
			return true
		}

		if slow == fast {
			break
		}
	}
	return false
}

func getSquaredSum(n int) int {
	res := 0
	for n > 0 {
		rem := n % 10
		n = int(n / 10)
		res += (rem * rem)
	}
	return res
}

func main() {
	sol := Solution{}
	fmt.Println(sol.find(23))
	fmt.Println(sol.find(12))
}
