package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	author = "ronething"
	github = "github.com/ronething"
	blog   = "https://blog.ronething.cn"
)

// 函数外面不能用 := 定义
// b := 1

func variableZeroValue() {
	var a int
	var s string
	// Printf %q 空串会把 "" 打出来
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "hello world"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c = 3, 4, true
	fmt.Println(a, b, c)
}

func variableShorter() {
	// := 定义赋值
	a, b, c, d := 3, 4, true, "不谈了"
	b = 5
	fmt.Println(a, b, c, d)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	//fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量的定义
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		python = iota
		java
		golang
	)
	fmt.Println(python, java, golang)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(author + " " + blog + " " + github)
	euler()
	triangle()
	consts()
	enums()
}
