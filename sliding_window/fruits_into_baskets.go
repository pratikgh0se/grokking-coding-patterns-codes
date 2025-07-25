package slidingwindow

import "fmt"

type Solution struct{}

func (s Solution) findLength(arr []string) int {

	basket := make(map[string]int)
	startPos := 0
	mostFruits := 0

	for endPos := 0; endPos < len(arr); endPos++ {

		currFruit := arr[endPos]
		if freq, has := basket[currFruit]; has {
			basket[currFruit] = freq + 1
		} else {
			basket[currFruit] = 1
		}

		for len(basket) > 2 {
			currFruit = arr[startPos]
			startPos++
			basket[currFruit]--
			if basket[currFruit] == 0 {
				delete(basket, currFruit)
			}
		}
		mostFruits = max(mostFruits, endPos-startPos+1)
	}
	return mostFruits
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	sol := Solution{}
	fmt.Println("Maximum number of fruits: ", sol.findLength([]rune{'A', 'B', 'C', 'A', 'C'}))
	fmt.Println("Maximum number of fruits: ", sol.findLength([]rune{'A', 'B', 'C', 'B', 'B', 'C'}))
}
