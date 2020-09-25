package main

import "strconv"

/**
	算法描述
	你正在和你的朋友玩 猜数字（Bulls and Cows）游戏：你写下一个数字让你的朋友猜。每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置都猜对了（称为“Bulls”, 公牛），有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋友将会根据提示继续猜，直到猜出秘密数字。

请写出一个根据秘密数字和朋友的猜测数返回提示的函数，用 A 表示公牛，用 B 表示奶牛。

请注意秘密数字和朋友的猜测数都可能含有重复数字。

示例 1:

输入: secret = "1807", guess = "7810"

输出: "1A3B"

解释: 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。
示例 2:

输入: secret = "1123", guess = "0111"

输出: "1A1B"

解释: 朋友猜测数中的第一个 1 是公牛，第二个或第三个 1 可被视为奶牛。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bulls-and-cows
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func getHint(secret string, guess string) string {
	//判断
	if len(secret) == 0 {
		return ""
	}

	//统计公牛数和母牛数
	Bulls,CoWs := 0, 0

	//声明hash计数
	hash := make(map[byte]int)

	for i := 0; i < len(secret);i ++ {
		hash[secret[i]] ++
	}

	//开始先公牛
	for i := 0 ; i < len(guess);i ++ {
		//判断是否为公牛
		if guess[i] == secret[i] {
			Bulls ++
			hash[guess[i]] --

		}
	}
	//猜母牛
	for i := 0 ; i < len(guess);i ++ {
		//判断是否为母牛
		if hash[guess[i]] > 0 {
			CoWs ++
			hash[guess[i]] --
		}
	}

	//返回
	return strconv.Itoa(Bulls) + "A" + strconv.Itoa(CoWs) + "B"
}