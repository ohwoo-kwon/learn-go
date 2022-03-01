package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}