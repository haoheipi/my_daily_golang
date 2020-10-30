package cn

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
// 示例：
//
// 给你这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
// 说明：
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// Related Topics 链表
// 👍 773 👎 0

// Time: 2020-10-24 13:22:02

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	// 给你这个链表：1->2->3->4->5
	//
	// 当 k = 3 时，应当返回: 3->2->1->4->5
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

//leetcode submit region end(Prohibit modification and deletion)
