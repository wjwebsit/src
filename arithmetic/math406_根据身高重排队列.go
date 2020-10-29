package main

import "fmt"

/**
	算法描述
	假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

注意：
总人数少于1100人。

示例

输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]

输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/queue-reconstruction-by-height
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reconstructQueue(people [][]int) [][]int {
	//获取总人数
	n := len(people)
	if n <= 1 {
		return people
	}

	//身高降序
	people = myInsetSort(people)

	//声明返回结果
	var result [][]int
	result = append(result,people[0])

	//构造结果
	for i := 1; i < len(people);i ++ {
		cur := people[i]

		index := 0
		j := 0
		for j < len(result) && index < cur[1] {
			if result[j][0] >= cur[0] {
				index++
			}
			j++
		}

		if j == len(result) {
			result = append(result,cur)
		} else {
			result = append(result,[]int{})
			k := len(result) - 1
			for k > j {
				result[k] = result[k - 1]
				k --
			}
			result[j] = cur
		}
	}

	//返回
	return result
}

/**
	根据前面有多少人进行插入排序--如果相等按身高降序排列
 */
func myInsetSort(p [][]int)[][]int {
	//开始排序
	for i := 1; i < len(p); i ++ {
		//获取当前
		key := p[i]

		//插入排序
		j := i - 1
		for j >= 0 && (p[j][0] < key[0] || (p[j][0] == key[0] && p[j][1] > key[1])) {
			p[j + 1] = p[j]
			j--
		}
		p[j + 1] = key
	}

	//返回
	return p
}
func main() {
	//声明身高
	people := [][]int{
		{7,0}, {4,4}, {7,1}, {5,0}, {6,1}, {5,2},
	}

	fmt.Println(reconstructQueue(people))
}