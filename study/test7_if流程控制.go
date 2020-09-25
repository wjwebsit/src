package main

import "fmt"

func main () { //左侧花括号 不允许下移

	fmt.Print(*Grade(85.4))

}

func Grade (score float64) *string {
	var str = new(string) //
	if score >= 90 {
		*str = "优秀"
	} else if score >= 80 {
		*str = "良"
	} else {
		*str = "及格"
	}

	return str

}
