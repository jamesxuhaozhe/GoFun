package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
func main() {
	fmt.Println(runtime.GOARCH)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}
