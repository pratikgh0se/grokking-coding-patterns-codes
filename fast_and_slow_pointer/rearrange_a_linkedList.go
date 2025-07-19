package fastandslowpointer

import "fmt"

//type ListNode struct {
//    Val  int
//    Next *ListNode
//}

type Solution struct{}

// reorder reorders the linked list as per the given logic
func (s *Solution) reorder(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var firstHalfLast *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		firstHalfLast = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	firstHalfLast.Next = nil
	reversed := reverse(slow)
	// TODO: Write your code here
	return merged(head, reversed)
}

func reverse(curr *ListNode) *ListNode {
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func merged(head1, head2 *ListNode) *ListNode {
	sol := head1

	for head1 != nil && head2 != nil {
		temp1 := head1.Next
		temp2 := head2.Next

		head1.Next = head2
		head2.Next = temp1

		head1 = temp1
		head2 = temp2
	}

	if head1 != nil {
		curr := sol
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = head1
	}

	if head2 != nil {
		curr := sol
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = head2
	}
	return sol
}

func main() {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 12}
	s := Solution{}
	s.reorder(head)
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
