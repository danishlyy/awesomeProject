package main

import "fmt"

func main() {
	var width,height,area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Println(area/10.0,"liters needed")
	width = 5.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0,"liters needed")
	fmt.Printf("area:%0.2f\n",1.0/3.0)
	rs := fmt.Sprintf("area:%0.2f\n",1.0/3.0)
	fmt.Printf(rs)
	fmt.Printf("A float:%f\n",3.1415)
	fmt.Printf("an integer:%d\n",14)
	fmt.Printf("a string:%s\n","hello")
	fmt.Printf("values:%v%v%v\n",1.2,"\t",true)
	fmt.Printf("values:%#v%#v%#v\n",1.2,"\t",true)
	fmt.Printf("types:%T %T %T\n",1.2,"\t",true)
	fmt.Printf("a boolean:%t\n",false)
	fmt.Printf("percent sign:%%\n")
	//sayHello()
}
