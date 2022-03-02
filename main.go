package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"yooa", "seunghee"}
	for _, person := range people {
		go isOMG(person, c)
	}
	result := <- c
	fmt.Println(result)
	fmt.Println(<- c)
}

func isOMG(person string, c chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	c <- true
}