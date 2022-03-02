package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [7]string{"hyojung", "mimi", "yooa", "seunghee", "jiho", "yubin", "arin"}
	for _, person := range people {
		go isOMG(person, c)
	}
	for i:=0;i<len(people);i++{
		fmt.Println(<-c)
	}
}

func isOMG(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- "Hello, " + person
}