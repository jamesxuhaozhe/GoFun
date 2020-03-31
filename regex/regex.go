package main

import (
	"fmt"
	"regexp"
)

const text = `My email address is jamesxuhaozhe@gmail.com
email is ssdfs@osgs.com
esdfsdf@cm.com.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
