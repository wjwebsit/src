package main

import "fmt"

/**
给定n个非负整数a 1，a 2， ...，a n，
每个数代表坐标中的一个点（i，  a i）。
在坐标内画n条垂直线，垂直线i  的两个端点分别为（i，  a i）和（i，0）。
找出其中的两条线，使得它们与  x  轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且  n  的值至少为2。

图中垂直线代表输入数组[1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
 */
func main() {
	//测试
	height := []int{1,8,6,2,5,4,8,3,7}

	fmt.Println(maxAreaTwoPointer(height))
}
/**
最开想到的
 */
func maxArea(height []int) int {
	//定义返回
	rtArea := 0

	//判断height数组的长度
	lenght := len(height)

	if lenght == 1 {
		return rtArea
	}

	if lenght == 2 {
		return getArea(1,height[0],2,height[1])
	}


	for i,_value := range height {

		for j := i + 1;j < lenght; j++ {
			tempArea := getArea(i,_value,j,height[j])

			if tempArea > rtArea {
				rtArea = tempArea
			}

		}

	}

	return rtArea



	//返回面积
	return rtArea
}
/**
获取面积xi,xh为第一个节点x和y
 */
func getArea(xi,xh,yi,yh int) int {
	if xh > yh {
		return (yi - xi) * yh
	} else {
		return (yi - xi) * xh
	}
}
/**
双指针法
这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。
此外，两线段距离越远，得到的面积就越大。

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。
此外，我们会使用变量 maxareamaxarea 来持续存储到目前为止所获得的最大面积。
在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxareamaxarea，并将指向较短线段的指针向较长线段那端移动一步。
 */
func maxAreaTwoPointer(height []int)int {
	//定义返回
	rtArea := 0

	//判断height数组的长度
	lenght := len(height)

	if lenght == 1 {
		return rtArea
	}

	if lenght == 2 {
		return getArea(1,height[0],2,height[1])
	}

	//定义起始和定义临时最左边和最右边
	beginIndex,endIndex := 0,len(height) - 1

	//循环获取
	for beginIndex < endIndex {
		tempArea := getArea(beginIndex + 1,height[beginIndex],endIndex + 1,height[endIndex])
		if tempArea > rtArea {
			rtArea = tempArea
		}
		if height[beginIndex] < height[endIndex] {
			beginIndex ++
		} else {
			endIndex --
		}

	}

	//返回
	return rtArea

}
