// 给一个分数，确定通过或者不通过
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("enter a grade:")
	// 控制台读取的都是字符串
	reader := bufio.NewReader(os.Stdin)
	for{
		input,err := reader.ReadString('\n')
		if err != nil || io.EOF == err{
			log.Fatal(err)
			break
		}
		input = strings.TrimSpace(input)
		result, err := strconv.ParseFloat(input, 64)
		var status string
		if result > 60 {
			status = "success"
			fmt.Println(status,"grade is ",result )
		}else {
			status = "fail"
			fmt.Println(status,"grade is ",result )
		}
	}



}