package main

import "sort"

/**
	在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
 */
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	//总面积
	sum := (C - A) * (D - B) + (G - E) * (H - F)

	//判断---不相交
	if E >= C || G <= A || F >= D || H <= B{
		return  sum
	}

	//获取相交矩形的宽
	tmp := []int{A,C,E,G}
	sort.Ints(tmp)
	w := tmp[2] - tmp[1]


	//同理获取gao
	tmp = []int{B,D,F,H}
	sort.Ints(tmp)
	h := tmp[2] - tmp[1]

	return sum - w * h

}