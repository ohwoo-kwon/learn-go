package main

import (
	"fmt"

	"github.com/ohwoo-kwon/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("ohwoo")
	account.Deposit(1000)
	account.ChangeOwner("junsang")
	fmt.Println(account)
	fmt.Println(account.Balance(), account.Onwer())
}