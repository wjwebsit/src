package main

import "fmt"

/**
	for ;;;{
	}
 */

func main() {

	fmt.Print(for1)
}

/**
	写法 1
 */
func for1 (n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}
/**
	while () {
	}
	类似while
 */
func for2(n int) int {
	var sum  = 0
	var i = 0

	for i <= n {
		sum += i
		i++
	}

	return sum

}


