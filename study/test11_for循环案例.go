package main

import "fmt"

func main() {
	juxing()
	fmt.Println("------------------")
	sanjiaoxing()
	fmt.Println("------------------")
	sanjiaoxing2()
	fmt.Println("------------------")
	sanJiaoXingRight()
	fmt.Println("------------------")
	dengyao()


}
//打印矩形5*6
func juxing() {
	for i := 1; i<=5;i++ {
		for j:=1;j<=6;j++ {
			fmt.Print("💗")
		}
		fmt.Println("")
	}
}

//打印直角三角形
func sanjiaoxing() {
	for i:= 1; i <= 5; i++ {
		for j:=1;j<=i;j++ {
			fmt.Print("💗")
		}
		fmt.Println("")
	}
}

func sanjiaoxing2() {
	for i:= 5; i >= 1; i-- {
		for j:=1;j<=i;j++ {
			fmt.Print("💗")
		}
		fmt.Println("")
	}
}
//右值三角形
func sanJiaoXingRight() {
	for i:=1;i<=5;i++ {
       	k := i
		for j:=1;j<=5;j++ {
       		if j <= 5 - k {
       			fmt.Print("  ")
			} else {
				fmt.Print("💗")
			}
	   	}
		fmt.Println(" ")
	}

}
//等腰三角形
func dengyao() {
	for i:=1;i<=5;i++ {
		k := i
		for j:=1;j<=5 + k -1;j++ {
			if j <= 5 - k {
				fmt.Print("  ")
			} else {
				fmt.Print("💗")
			}
		}
		fmt.Println(" ")
	}
}
