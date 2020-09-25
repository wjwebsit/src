package main

/**
	算法描述
	示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//利用辅助数组
	var arr []int
	p := head
	for  p != nil {
		arr = append(arr,p.Val)
		p = p.Next
	}

	//回文
	s,e := 0 , len(arr) - 1

	for s < e && arr[s] == arr[e] {
		s ++
		e --
	}

	return s >= e
}