package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
	算法描述：（略）
	示例：

输入: n = 1
返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
 

提示：

输出的顺序没有要求。
小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
超过表示范围（小时 0-11，分钟 0-59）的数据将会被舍弃，也就是说不会出现 "13:00", "0:61" 等时间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-watch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//定义全局预处理
var H  = make(map[int][]string)

var I = make(map[int][]string)

func initTime() {
	//预处理时分
	if len(H) > 0 && len(I) > 0 {
		return
	}

	//获取时针
	H[0] = []string{"0"}
	for i := 1; i <= 3; i ++ {
		getTime([]int{1,2,4,8},0,0,i,0,11,&H)
		sort.Slice(H[i],func(a,b int)bool {
			num1 , _ := strconv.Atoi(H[i][a])
			num2 , _ := strconv.Atoi(H[i][b])
			return  num1 < num2
		})

	}



	//分针
	I[0] = []string{"00"}
	for i := 1; i <= 5; i ++ {
		getTime([]int{1,2,4,8,16,32},0,0,i,0,59,&I)
		sort.Slice(I[i],func(a,b int)bool {
			num1 , _ := strconv.Atoi(I[i][a])
			num2 , _ := strconv.Atoi(I[i][b])
			return  num1 < num2
		})
	}

}
func getTime(times []int,s int, cur int,n int ,sum int, limit int ,result *map[int][]string) {
	//判断
	if cur == n {
		//判断
		if sum <= limit {
			//写入结果
			if sum < 10 && len(times) > 4 {
				(*result)[n] = append((*result)[n],"0" +strconv.Itoa(sum))
			} else {
				(*result)[n] = append((*result)[n],strconv.Itoa(sum))
			}
		}

		return
	}

	//写入
	for i := s; i < len(times); i ++ {
		getTime(times,i + 1,cur + 1,n,sum + times[i],limit,result)
	}


}


func readBinaryWatch(num int) []string {
	//判断特殊情况
	if num > 8 {
		return []string{}
	}
	if num == 0 {
		return []string{"0:00"}
	}

	//结果
	var result []string

	//预处理
	initTime()

	//时针最多只能亮3只
	for i := 0;i <= 3 && i <= num; i ++ {
		//分针
		j := num - i

		//判断
		if j > 5 {
			continue
		}

		//构造
		for _,h := range H[i] {
			for _,k := range I[j] {
				result = append(result,h + ":" + k)
			}
		}


	}
	return result

}

func main() {
	fmt.Println(readBinaryWatch(1))
}