package main

import (
	"fmt"
	"log"

	"github.com/ohwoo-kwon/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}