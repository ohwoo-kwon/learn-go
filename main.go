package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"hyojung":"leader"}
	baseWord := "mimi"
	dictionary.Add(baseWord, "rapper")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	definition, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}