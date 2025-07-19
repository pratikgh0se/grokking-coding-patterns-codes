import (
	"math"
)

type Solution struct{}

func (s *Solution) sort(arr []int) int {
	// TODO: Write your code here
	low, high, n := 0, len(arr)-1, len(arr)-1
	for low < n && arr[low] <= arr[low+1] {
		low++
	}
	if low == n {
		return 0
	}
	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}
	min, max := findMinMax(arr, low, high)
	for low > 0 && min < arr[low-1] {
		low--
	}
	for high < n && max > arr[high+1] {
		high++
	}
	return high - low + 1
}

func findMinMax(arr []int, l, r int) (int, int) {
	min, max := math.MaxInt32, math.MinInt32
	for i := l; i <= r; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

func main() {
	sol := Solution{}
	fmt.Println(sol.sort([]int{1, 2, 5, 3, 7, 10, 9, 12}))
	fmt.Println(sol.sort([]int{1, 3, 2, 0, -1, 7, 10}))
	fmt.Println(sol.sort([]int{1, 2, 3}))
	fmt.Println(sol.sort([]int{3, 2, 1}))
}