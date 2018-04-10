package main

import (
	"fmt"
)

func doSomthing() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ) Hello men!", i)
	}
}

func main() {
	fmt.Println("Привет. Мир!")
	fmt.Println("Привет. Друзья!")
	fmt.Println("Привет. Люди!")
	doSomthing()
}
