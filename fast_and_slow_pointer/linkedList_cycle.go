package fastandslowpointer

import "fmt"

// ListNode struct represents a node in the linked list
// type ListNode struct {
//	Val int
//	Next  *ListNode
// }

// Solution struct is used to group the methods
type Solution struct{}

// hasCycle method checks for a cycle in the linked list
// It keeps the same logic and comments as the Java version
func (s *Solution) hasCycle(head *ListNode) bool {
	// TODO: Write your code here
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	sol := Solution{}
	fmt.Println("LinkedList has cycle:", sol.hasCycle(head))

	head.Next.Next.Next.Next.Next.Next = head.Next.Next
	fmt.Println("LinkedList has cycle:", sol.hasCycle(head))

	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next
	fmt.Println("LinkedList has cycle:", sol.hasCycle(head))
}
