package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我喜欢中国!" // utf-8
	//fmt.Println(len(s))
	fmt.Println(s)
	fmt.Printf("%X\n", []byte(s))

	for _, byte := range []byte(s) {
		fmt.Printf("%X ", byte)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		//fmt.Println("size ", size)
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
