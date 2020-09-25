package main
/**
	算法描述
	珠玑妙算游戏（the game of master mind）的玩法如下。

计算机有4个槽，每个槽放一个球，颜色可能是红色（R）、黄色（Y）、绿色（G）或蓝色（B）。例如，计算机可能有RGGB 4种（槽1为红色，槽2、3为绿色，槽4为蓝色）。作为用户，你试图猜出颜色组合。打个比方，你可能会猜YRGB。要是猜对某个槽的颜色，则算一次“猜中”；要是只猜对颜色但槽位猜错了，则算一次“伪猜中”。注意，“猜中”不能算入“伪猜中”。

给定一种颜色组合solution和一个猜测guess，编写一个方法，返回猜中和伪猜中的次数answer，其中answer[0]为猜中的次数，answer[1]为伪猜中的次数。

示例：

输入： solution="RGBY",guess="GGRR"
输出： [1,1]
解释： 猜中1次，伪猜中1次。
提示：

len(solution) = len(guess) = 4
solution和guess仅包含"R","G","B","Y"这4种字符

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/master-mind-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func masterMind(solution string, guess string) []int {
	//统计猜中
	T := 0

	//采用hash
	sHash,gHash := make(map[byte]int),make(map[byte]int)
	for i := 0 ; i < 4;i ++ {
		if solution[i] == guess[i] {
			T ++
		} else {
			//槽中字符
			sHash[solution[i]]++

			//猜的字符
			gHash[guess[i]]++
		}
	}

	//统计伪猜中
	F := 0
	if gHash != nil {
		for k, v := range gHash {
			//判断
			if sHash[k] > 0 {//表示存在
				F += minAb(v,sHash[k])
			}
		}

	}

	//返回
	return []int{T,F}

}

func minAb (a,b int)int {
	if a < b {
		return  a
	}
	return b
}