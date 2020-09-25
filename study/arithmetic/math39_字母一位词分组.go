package main

import "fmt"

/**
	算法描述:
	给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
 */

 /**
 	算法要点 adc acd dcd   "a"=>质数1 "d "=> 质数2 "d"=> 质数3  只要积相同 则为一组（秒）
  */
func groupAnagrams(strs []string) [][]string {
	//声明结果
	var result [][]string

	//判断
	if len(strs) == 0 {
		return result
	}

	//声明字母质数
	var mp = map[byte]int{'a':2,'b':3,'c':5,'d':7,'e':11,'f':13,'g':17,'h':19,'i':23,'j':29,'k':31,'l':37,'m':41,'n':43,'o':47,'p':53,'q':59,'r':61, 's':67, 't':71, 'u':73, 'v':79, 'w':83, 'x':89, 'y':97, 'z':101}

	//team
	team := map[int][]string{}

	//分组
	for i := 0; i < len(strs); i ++ {
		sum := 1
		for j := 0; j < len(strs[i]);j ++ {
			sum *= mp[strs[i][j]]
		}

		//存
		team[sum] = append(team[sum],strs[i])
	}


	//处理结果
	for _,v := range team {
		result = append(result,v)
	}

	//返回结果
	return result

}

//测试
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
