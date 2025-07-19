package fastandslowpointer

import "fmt"

// type ListNode struct {
// 	Val int
// 	Next  *ListNode
// }

type Solution struct{}

func (s *Solution) isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// TODO: Write your code here
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	reversed := reverse(slow)
	for head != nil && reversed != nil {
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	sol := Solution{}
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	fmt.Printf("Is palindrome: %v\n", sol.isPalindrome(head))

	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	fmt.Printf("Is palindrome: %v\n", sol.isPalindrome(head))
}
