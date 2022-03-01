package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"hyojung":"leader"}
	definition, err := dictionary.Serach("hyojung")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	fmt.Println(dictionary)
}