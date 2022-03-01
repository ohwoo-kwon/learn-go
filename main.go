package main

import "fmt"

func main() {
	ohmygirl := [7] string{"hyojung", "mimi", "yooa"}
	ohmygirl[3] = "seunghee"
	ohmygirl[4] = "jiho"
	ohmygirl[5] = "yubin"
	ohmygirl[6] = "arin"
	fmt.Println(ohmygirl)

	names := []string{"ohwoo"}
	names = append(names, "anonymous1")
	names = append(names, "anonymous1")
	fmt.Println(names)
}