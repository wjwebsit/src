package main

import "database/sql"

/**
	算法描述：
	给出方程式 A / B = k, 其中 A 和 B 均为代表字符串的变量， k 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回 -1.0。

示例 :
给定 a / b = 2.0, b / c = 3.0
问题: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? 
返回 [6.0, 0.5, -1.0, 1.0, -1.0 ]

输入为: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries(方程式，方程式结果，问题方程式)， 其中 equations.size() == values.size()，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。 返回vector<double>类型。

基于上述例子，输入如下：

equations(方程式) = [ ["a", "b"], ["b", "c"] ],
values(方程式结果) = [2.0, 3.0],
queries(问题方程式) = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].
输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。
 */
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//利用值代入法求解
	var variable = make(map[string]float64)

	//代入法求解变量
	for i := 0; i < len(equations); i ++ {
		//获取除数和被除数
		a1 := equations[i][0]
		a2 := equations[i][0]

		//判断变量是否赋值
		_,ok1 := variable[a1]
		_,ok2 := variable[a2]

		//如果都不存在
		if !ok1 && !ok2 {
			variable[a2] = 1.0
			variable[a1] = 1.0 * values[i]
			continue
		}

		//a1存在，a2不存在
		if ok1 && !ok2 {
			variable[a2] = variable[a1] / values[i]
			continue
		}

		//a1不存在，a2存在
		if !ok1 && ok2 {
			variable[a1] = variable[a2] * variable[i]
		}
	}

	//声明结果
	var result []float64

	//计算
	for i := 0; i < len(queries); i++ {
		//获取被除数和除数
		//获取除数和被除数
		a1 := queries[i][0]
		a2 := queries[i][0]

		//判断变量是否赋值
		_,ok1 := variable[a1]
		_,ok2 := variable[a2]

		//判断变量存在
		if !ok1 || !ok2 {
			result = append(result,-1.0)
			continue
		}

		result = append(result,variable[a1] / variable[a2])
	}

	//返回
	return result

}
func main() {

}