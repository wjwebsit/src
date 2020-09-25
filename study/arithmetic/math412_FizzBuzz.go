package main

import "strconv"

/**
	算法描述
	写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

示例：

n = 15,

返回:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
 */
func fizzBuzz(n int) []string {
	//定义
	f := 3
	z := 5
	fz := 15

	//定义结果集
	var result []string

	for i := 1; i <= n ; i ++ {
		if i == fz {
			result = append(result,"FizzBuzz")
			f += 3
			z += 5
			fz += 15
		} else if i == f {
			result = append(result,"Fizz")
			f += 3
		} else if i == z {
			result = append(result,"Buzz")
			z += 5
		} else {
			result = append(result,strconv.Itoa(i))
		}

	}

	//返回
	return result

}