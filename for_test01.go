package main

import "fmt"

func main() {
	for x:=3;x<10;x++{
		fmt.Print(x ,"   ")
	}

	for x:=3;x<10;x++{
		x += 1
		fmt.Print(x ,"   ")
	}
}
