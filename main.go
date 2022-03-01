package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"hyojung":"leader"}
	word := "mimi"
	definition := "rapper"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	mimi, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", mimi)
	newErr := dictionary.Add(word, definition)
	if newErr != nil {
		fmt.Println(newErr)
	}
}