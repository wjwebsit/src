package main

import "fmt"

/**

 */

func maxSizeSlices(slices []int) int {


}

func myMaxSizeSlices(slice []int,s,e int,dp *[520][520]int) int{
	//判断
	if s + 2 == e {
		(*dp)[s][e] = slice[s] + slice[s + 1] + slice[s + 2]
		return (*dp)[s][e]

	}

	//判断缓存
	if (*dp)[s][e] > 0 {
		return (*dp)[s][e]
	}
	max := 0
	//循环
	for i := s; i <= e ; i ++ {
		if i == s {
			max = maxAB2(max,slice[i] + slice[i + 1] + slice[e] + myMaxSizeSlices(slice,i + 2,e - 1,dp))

		} else if i == e {
			max = maxAB2(max,slice[i] + slice[s] + slice[i - 1] + myMaxSizeSlices(slice,s + 1,e - 2,dp))

		}else {
			max = maxAB2(max,slice[i] + slice[i + 1] + slice[i - 1] + myMaxSizeSlices(slice,i + 2,e - 1,dp))

		}


	}

}
func maxAB2(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	slice := []int{1,2,3,4,5,6}
	fmt.Println(maxSizeSlices(slice))
}