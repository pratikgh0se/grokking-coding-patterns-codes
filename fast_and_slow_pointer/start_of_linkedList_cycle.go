package fastandslowpointer

import "fmt"

// ListNode struct definition
// type ListNode struct {
// 	Val int
// 	Next  *ListNode
// }

// Solution struct definition
type Solution struct{}

// findCycleStart method of Solution struct
func (s *Solution) findCycleStart(head *ListNode) *ListNode {
	// TODO Write your code here
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			break
		}
	}
	return slow
}

// main function
func main() {
	sol := Solution{}
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	// Create a cycle by connecting nodes
	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	fmt.Println("LinkedList cycle start: ", sol.findCycleStart(head).Val)

	// Create a different cycle
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next
	fmt.Println("LinkedList cycle start: ", sol.findCycleStart(head).Val)

	// Create a cycle that points back to the head
	head.Next.Next.Next.Next.Next.Next = head
	fmt.Println("LinkedList cycle start: ", sol.findCycleStart(head).Val)
}
