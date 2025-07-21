type Solution struct{}

func (s *Solution) loopExists(arr []int) bool {
	// TODO: Write your code here
	n := len(arr)
	visitedForward := make([]bool, n)
	visitedBackward := make([]bool, n)

	for i, _ := range arr {
		nextIdx := getNextIndex(i, arr[i], n)
		if nextIdx == i {
			continue
		}
		slow, fast := i, i
		currForwardOrNot := isFoward(arr[i])
		if currForwardOrNot {
			if visitedForward[i] {
				continue
			}
		} else {
			if visitedBackward[i] {
				continue
			}
		}
		for slow < n && fast < n {
			if currForwardOrNot {
				visitedForward[slow] = true
				visitedForward[fast] = true
			} else {
				visitedBackward[slow] = true
				visitedBackward[fast] = true
			}

			slow = getNextIndex(slow, arr[slow], n)
			if currForwardOrNot != isFoward(arr[slow]) {
				break
			}

			fastNext := getNextIndex(fast, arr[fast], n)
			if currForwardOrNot != isFoward(arr[fastNext]) {
				break
			}

			if currForwardOrNot {
				visitedForward[fastNext] = true
			} else {
				visitedBackward[fastNext] = true
			}

			fast = getNextIndex(fastNext, arr[fastNext], n)
			if currForwardOrNot != isFoward(arr[fast]) {
				break
			}

			if slow == fast {
				nextPos := getNextIndex(slow, arr[slow], n)
				if nextPos == slow {
					break
				}
				return true
			}

		}
	}
	return false
}

func getNextIndex(currIdx, jump, n int) int {
	return (currIdx + (jump % n) + n) % n
}

func isFoward(i int) bool {
	if i < 0 {
		return false
	}
	return true
}
func main() {
	sol := Solution{}
	fmt.Println(sol.loopExists([]int{1, 2, -1, 2, 2}))
	fmt.Println(sol.loopExists([]int{2, 2, -1, 2}))
	fmt.Println(sol.loopExists([]int{2, 1, -1, -2}))
}
