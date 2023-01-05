package main

import "fmt"

func someFunc() (someInt int) {
	fmt.Println("calling someFunc from utils")
	return 123
}
