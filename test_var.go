package main

import "fmt"

func main() {
	a := 1
	// 只要短变量声明中至少有一个变量名是新的，这是允许的。新变量名被视为声明，而现有的名字被视为赋值
	b,a := 2,3
	a,c := 4,5
	fmt.Println(a,b,c)

}
