package main
/**
go 语言的传引用一般通过指针来实现的------递归函数的心得核心也是基于传引用实现
 */
import "fmt"

func main() {
	a := 20

	changeValue(a)
	fmt.Println(a)

	changeRefValue(&a)
	fmt.Println(a)
}

func changeValue(a int) {
	a = 10
}

func changeRefValue(a *int) {
	*a = 10
}
