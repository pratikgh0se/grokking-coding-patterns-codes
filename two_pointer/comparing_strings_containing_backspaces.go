package twopointer

import "fmt"

// Solution struct simulates a class in Go.
type Solution struct{}

// compare is a method of Solution struct.
// It uses two pointer approach to compare the strings.
func (s *Solution) compare(str1, str2 string) bool {
	// TODO: Write your code here
	idx1 := len(str1) - 1
	idx2 := len(str2) - 1

	for idx1 >= 0 && idx2 >= 0 {
		p1 := getNextValidIndex(str1, idx1)
		p2 := getNextValidIndex(str2, idx2)
		if p1 == -1 && p2 == -1 {
			return true
		}
		if p1 == -1 || p2 == -1 {
			return false
		}
		if str1[p1] != str2[p2] {
			return false
		}
		idx1 = p1 - 1
		idx2 = p2 - 1
	}
	return true
}

func getNextValidIndex(s string, p int) int {
	backSpaceCount := 0
	for p >= 0 {
		if s[p] == '#' {
			backSpaceCount++
			p--
		} else if backSpaceCount > 0 {
			backSpaceCount--
			p--
		} else {
			return p
		}
	}
	return -1

}

func main() {
	solution := Solution{}
	fmt.Println(solution.compare("xy#z", "xzz#"))
	fmt.Println(solution.compare("xy#z", "xyz#"))
	fmt.Println(solution.compare("xp#", "xyz##"))
	fmt.Println(solution.compare("xywrrmp", "xywrrmu#p"))
}
