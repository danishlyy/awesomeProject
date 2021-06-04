package main

import "fmt"

func sayHello() {
	fmt.Println("hello,world!")
}

func repeatLine(line string, times int) {
	for i:=0;i<times;i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width int64,height int64)  {
	result := width * height
	fmt.Println("area is ", result/1.0)
}

func main() {
	sayHello()
	repeatLine("hello",3)
	paintNeeded(10,20)
}