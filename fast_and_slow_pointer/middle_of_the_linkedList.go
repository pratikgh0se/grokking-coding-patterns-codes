package fastandslowpointer

import "fmt"

// type ListNode struct {
// 	Val int
// 	Next  *ListNode
// }

type Solution struct{}

// findMiddle finds the middle node in a linked list
func (s *Solution) findMiddle(head *ListNode) *ListNode {
	// TODO: Write your code here
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
	}
	return head
}

func main() {
	// Create a linked list with nodes containing values 1, 2, 3, 4, and 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Create a Solution instance
	s := Solution{}

	// Find and print the middle node of the linked list
	fmt.Println("Middle Node:", s.findMiddle(head).Val)

	// Add more nodes to the linked list
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	// Find and print the middle node again
	fmt.Println("Middle Node:", s.findMiddle(head).Val)

	// Add another node to the linked list
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}

	// Find and print the middle node once more
	fmt.Println("Middle Node:", s.findMiddle(head).Val)
}
