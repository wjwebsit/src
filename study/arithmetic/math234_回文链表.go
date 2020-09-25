package main
/**
	算法描述
	请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 */
func isPalindrome(head *ListNode) bool {
	//判断
	if head == nil {
		return false
	}

	//声明辅助数组
	var nums []int
	p := head
	for p != nil {
		nums = append(nums,p.Val)
		p = p.Next
	}

	//判断回文
	s,e := 0,len(nums) - 1
	for s <= e && nums[s] == nums[e] {
		s++
		e--
	}

	//返回
	return  s >= e
}