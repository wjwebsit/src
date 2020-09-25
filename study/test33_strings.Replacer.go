package main

import (
	"fmt"
	"strings"
)

func main() {
	//将所有 a => A b => B
	relacer := strings.NewReplacer("a","A","b","B")

	test := relacer.Replace("haha Bob")

	fmt.Println(test) //hAhA BoB

}
