package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false
	// 把时间转换成Unix时间格式，这是一个整数，是自1970年1月1日以来的秒数
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("target number is : ",target)
	fmt.Println("please enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	for guesses:=0;guesses<10;guesses++ {
		fmt.Println("you have ",10-guesses, " guesses left")
		input,err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str := strings.TrimSpace(input)
		num, numberErr := strconv.Atoi(str)
		if numberErr != nil{
			log.Fatal(numberErr)
		}
		if num < target {
			fmt.Println("you guess low,your number is",num)
		}else if num > target {
			fmt.Println("you guess high,your number is",num)
		}else {
			success = true
			fmt.Println("you guess number is equal target")
			break
		}
	}

	if !success {
		fmt.Println("sorry,you cannot guess my number,it was ",target)
	}

}
