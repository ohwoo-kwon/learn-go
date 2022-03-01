package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(1000)
	fmt.Println(account.Balance())
}