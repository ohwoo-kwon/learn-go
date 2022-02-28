package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upperName := lenAndUpper("ohwoo")
	fmt.Println(totalLength, upperName)
	repeatMe("hyojung", "mimi", "yooa", "seunghee", "jiho", "yubin", "arin")
}