package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("hyojung")
	go sayHello("mimi")
	time.Sleep(time.Second * 3)
}

func sayHello(person string){
	for i := 0; i < 10; i++ {
		fmt.Println(person, "Hello!", i)
		time.Sleep(time.Second)
	}
}