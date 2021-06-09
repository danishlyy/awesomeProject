package main

import (
	"errors"
	"fmt"
	"math"
)

var packageVariable int64

func sayHello() {
	fmt.Println("hello,world!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width int64, height int64) {
	result := width * height
	fmt.Println("area is ", result/1.0)
}

func paintNeededReturn(width float64, height float64) float64 {
	area := width * height
	return area / 1.0
}

// 函数返回值
func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 0 {
		err := errors.New("grade should bigger than zero")
		return err.Error()
	}
	if grade < 60 {
		return "failing"
	}
	return "passing"
}

// 返回值为多个
func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cannot get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func main() {
	root, err := squareRoot(-9.3)
	if err != nil {
		// cannot get square root of negative number
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f", root)
	}
	cans, remainder := floatParts(1.26)
	// 1 0.26
	fmt.Println(cans, remainder)
	// grade should bigger than zero
	fmt.Println(status(-10))
	var amount, total float64
	amount = paintNeededReturn(4.2, 3.0)
	// 12.60 liters needed
	fmt.Printf("%0.2f liters needed\n\n", amount)
	total += amount
	amount = paintNeededReturn(5.2, 3.5)
	// 18.20 liters needed
	fmt.Printf("%0.2f liters needed\n\n", amount)
	total += amount
	// total:30.80 liters
	fmt.Printf("total:%0.2f liters\n", total)
	grade := status(80)
	// passing
	fmt.Println(grade)
	dozen := double(6.0)
	// 12
	fmt.Println(dozen)
	// 8.4
	fmt.Println(double(4.2))
	packageVariable = 10
	fmt.Println(packageVariable)
	sayHello()
	repeatLine("hello", 3)
	paintNeeded(10, 20)
}
