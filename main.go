package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "chickenn"}
	nico := person{name: "nico", age:18, favFood:favFood}
	fmt.Println(nico)
}