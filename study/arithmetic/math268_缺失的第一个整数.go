package main
/**
	//算法描述；
 */
func missingNumber(nums []int) int {
	//获取n
	/*n := len(nums)

	mp := make(map[int]int)
	s, e := 0, n-1

	for s <= e {
		mp[nums[s]] = 1
		mp[nums[e]] = 1
		s++
		e--
	}
	s, e = 0, n
	for s <= e {
		if mp[s] == 0 {
			return s
		}
		if mp[e] == 0 {
			return e
		}
		s++
		e--
	}
	return - 1*/
	//采用高斯定理
	n := len(nums)
	sum := (1 + n) * (n) / 2
	for i := 0;i < n;i++ {
		sum -= nums[i]
	}
	return  sum

}