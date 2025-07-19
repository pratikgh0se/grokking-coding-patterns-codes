type Solution struct{}

func (s *Solution) loopExists(arr []int) bool {
	// TODO: Write your code here
	n := len(arr)
	visited := make([]bool, n)

	for i, _ := range arr {
		if visited[i] {
			continue
		}
		nextIdx := getNextIndex(i, arr[i], n)
		if nextIdx == i {
			continue
		}
		slow, fast := i, i
		currForwardOrNot := isFoward(arr[i])
		for slow < n && fast < n {
			visited[slow] = true
			visited[fast] = true
			slow = getNextIndex(slow, arr[slow], n)
			if currForwardOrNot != isFoward(arr[slow]) {
				break
			}
			fastNext := getNextIndex(fast, arr[fast], n)
			if currForwardOrNot != isFoward(arr[fastNext]) {
				break
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
