/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// tmp -> node1 -> node2
// 1 添加哑结点,方便统一处理 nil 的情况
// 记下当前节点后面两个节点的地址,要不然就找不到了.
// 顺序是先把当前处理的节点的 next 指向 2 号节点,
// 然后把 下次要处理的节点 记到 node1 里面
// 然后那 node1 接到 node2 后面
// 然后把 tmp 指向 node ,下一轮 处理 后面两个,如果都不是 nil 的话.
func swapPairs(head *ListNode) *ListNode {
	dummyhead := &ListNode{0, head}
	tmp := dummyhead
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}
	return dummyhead.Next
}

// @lc code=end

